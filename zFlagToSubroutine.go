package netroutine

import (
	"context"
	"encoding/json"
)

const idFlagToSubroutine = "FlagToSubroutine"

func init() {
	blocks[idFlagToSubroutine] = &FlagToSubroutine{}
}

type FlagToSubroutine struct {
	FromKey string
	IfTrue  Routine
	IfFalse Routine
}

func (b *FlagToSubroutine) toBytes() ([]byte, error) {
	var t struct {
		FromKey string
		IfTrue  []byte
		IfFalse []byte
	}

	t.FromKey = b.FromKey

	ifTrue, err := b.IfTrue.ToBytes()
	if err != nil {
		return nil, err
	}

	t.IfTrue = ifTrue

	ifFalse, err := b.IfFalse.ToBytes()
	if err != nil {
		return nil, err
	}

	t.IfFalse = ifFalse

	return json.Marshal(t)
}

func (b *FlagToSubroutine) fromBytes(bytes []byte) error {
	var t struct {
		FromKey string
		IfTrue  []byte
		IfFalse []byte
	}

	err := json.Unmarshal(bytes, &t)
	if err != nil {
		return err
	}

	ifTrue, err := RoutineFromBytes(t.IfTrue)
	if err != nil {
		return err
	}

	b.IfTrue = *ifTrue

	ifFalse, err := RoutineFromBytes(t.IfFalse)
	if err != nil {
		return err
	}

	b.IfFalse = *ifFalse

	return nil
}

func (b *FlagToSubroutine) kind() string {
	return idFlagToSubroutine
}

func (b *FlagToSubroutine) Run(ctx context.Context, wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	bo, ok := v.(bool)
	if !ok {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	if bo {
		wce.Run(ctx, &b.IfTrue)
		return log(b, "ran true routine", wce.Status)
	} else {
		wce.Run(ctx, &b.IfFalse)
		return log(b, "ran false routine", wce.Status)
	}
}
