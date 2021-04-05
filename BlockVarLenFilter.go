package netroutine

import (
	"encoding/json"
)

const idBlockVarLenFilter = "BlockVarLenFilter"

type BlockVarLenFilter struct {
	FromKey string
	Length  int
}

func (b *BlockVarLenFilter) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockVarLenFilter) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockVarLenFilter) kind() string {
	return idBlockVarLenFilter
}

func (b *BlockVarLenFilter) Run(wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "couldn't find the variable", Error)
	}

	s, err := toString(v)
	if err != nil {
		return log(b, "unable to convert variable to string", Error)
	}

	if len(s) < b.Length {
		return log(b, "variable not long enough", Fail)
	}
	return log(b, "variable long enough", Success)
}
