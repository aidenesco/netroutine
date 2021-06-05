package netroutine

import (
	"context"
	"encoding/json"
)

func init() {
	blocks[idSliceToSubroutine] = &SliceToSubroutine{}
}

const idSliceToSubroutine = "SliceToSubroutine"

type SliceToSubroutine struct {
	VariableKeys []string
	Routine      Routine
	FromKeys     []string
	IgnoreErrors bool
}

func (b *SliceToSubroutine) toBytes() ([]byte, error) {
	var t struct {
		VariableKeys []string
		Routine      []byte
		FromKeys     []string
		IgnoreErrors bool
	}

	t.VariableKeys = b.VariableKeys
	t.FromKeys = b.FromKeys
	t.IgnoreErrors = b.IgnoreErrors

	by, err := b.Routine.ToBytes()
	if err != nil {
		return nil, err
	}

	t.Routine = by

	return json.Marshal(t)
}

func (b *SliceToSubroutine) fromBytes(bytes []byte) error {
	var t struct {
		VariableKeys []string
		Routine      []byte
		FromKeys     []string
		IgnoreErrors bool
	}

	err := json.Unmarshal(bytes, &t)
	if err != nil {
		return err
	}

	b.VariableKeys = t.VariableKeys
	b.FromKeys = t.FromKeys
	b.IgnoreErrors = t.IgnoreErrors

	routine, err := RoutineFromBytes(t.Routine)
	if err != nil {
		return err
	}

	b.Routine = *routine

	return nil
}

func (b *SliceToSubroutine) kind() string {
	return idSliceToSubroutine
}

func (b *SliceToSubroutine) Run(ctx context.Context, wce *Environment) (string, Status) {
	var sources [][]interface{}
	var length int

	for _, v := range b.FromKeys {
		data, ok := wce.getData(v)
		if !ok {
			return log(b, missingWorkingData(v), Error)
		}

		sl, ok := data.([]interface{})
		if !ok {
			return log(b, reportWrongType(v), Error)
		}

		sources = append(sources, sl)
	}

	if len(sources) == 0 {
		return log(b, "no sources found", Error)
	}

	length = len(sources[0])

	for i := 1; i < len(sources); i++ {
		if len(sources[i]) != length {
			return log(b, "got slices of varying lengths", Error)
		}
	}

	if len(sources) != len(b.VariableKeys) {
		return log(b, "got variable keys and sources of different lengths", Error)
	}

	for i := 0; i < length; i++ {
		for j, v := range sources {
			wce.setData(b.VariableKeys[j], v[i])
		}

		wce.Run(ctx, &b.Routine)

		if !b.IgnoreErrors && wce.Status != Success {
			return log(b, "got a status other than Success", wce.Status)
		}
	}

	return log(b, "ran subroutines", wce.Status)
}
