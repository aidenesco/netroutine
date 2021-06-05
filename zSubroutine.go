package netroutine

import (
	"context"
	"encoding/json"
)

const idSubroutine = "Subroutine"

func init() {
	blocks[idSubroutine] = &Subroutine{}
}

type Subroutine struct {
	Subroutine Routine
}

func (b *Subroutine) toBytes() ([]byte, error) {
	var t struct {
		Subroutine []byte
	}

	subroutineBytes, err := b.Subroutine.ToBytes()
	if err != nil {
		return nil, err
	}

	t.Subroutine = subroutineBytes

	return json.Marshal(t)
}

func (b *Subroutine) fromBytes(bytes []byte) error {
	var t struct {
		Subroutine []byte
	}

	err := json.Unmarshal(bytes, &t)
	if err != nil {
		return err
	}

	routine, err := RoutineFromBytes(t.Subroutine)
	if err != nil {
		return err
	}

	b.Subroutine = *routine

	return nil
}

func (b *Subroutine) kind() string {
	return idSubroutine
}

func (b *Subroutine) Run(ctx context.Context, wce *Environment) (string, Status) {

	wce.Run(ctx, &b.Subroutine)

	return log(b, "ran subroutine", wce.Status)
}
