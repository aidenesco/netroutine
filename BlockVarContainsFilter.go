package netroutine

import (
	"encoding/json"
	"strings"
)

const idBlockVarContainsFilter = "BlockVarContainsFilter"

type BlockVarContainsFilter struct {
	OneOf   []string
	FromKey string
}

func (b *BlockVarContainsFilter) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockVarContainsFilter) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockVarContainsFilter) kind() string {
	return idBlockVarContainsFilter
}

func (b *BlockVarContainsFilter) Run(wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "couldn't find the variable", Error)
	}

	s, err := toString(v)
	if err != nil {
		return log(b, "unable to convert variable to string", Error)
	}

	for _, v := range b.OneOf {
		if strings.Contains(s, v) {
			return log(b, "found contained string:"+v, Success)
		}
	}

	return log(b, "variable doesnt contain any of the provided strings", Fail)
}
