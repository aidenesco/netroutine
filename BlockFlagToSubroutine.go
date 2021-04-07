package netroutine

import (
	"encoding/json"
)

const idBlockFlagToSubroutine = "BlockFlagToSubroutine"

type BlockFlagToSubroutine struct {
	FromKey string
	IfTrue  Routine
	IfFalse Routine
}

func (b *BlockFlagToSubroutine) toBytes() ([]byte, error) {
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

func (b *BlockFlagToSubroutine) fromBytes(bytes []byte) error {
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

func (b *BlockFlagToSubroutine) kind() string {
	return idBlockFlagToSubroutine
}

func (b *BlockFlagToSubroutine) Run(wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "variable not found", Error)
	}

	bo, ok := v.(bool)
	if !ok {
		return log(b, "variable not a boolean", Error)
	}

	if bo {
		b.IfTrue.Run(wce)
		return log(b, "ran true routine", wce.Status)
	} else {
		b.IfFalse.Run(wce)
		return log(b, "ran false routine", wce.Status)
	}
}
