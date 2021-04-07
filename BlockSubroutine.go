package netroutine

import (
	"encoding/json"
)

const idBlockSubroutine = "Subroutine"

type BlockSubroutine struct {
	Subroutine Routine
}

func (b *BlockSubroutine) toBytes() ([]byte, error) {
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

func (b *BlockSubroutine) fromBytes(bytes []byte) error {
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

func (b *BlockSubroutine) kind() string {
	return idBlockSubroutine
}

func (b *BlockSubroutine) Run(wce *Environment) (string, Status) {

	b.Subroutine.Run(wce)

	return log(b, "ran subroutine", wce.Status)
}
