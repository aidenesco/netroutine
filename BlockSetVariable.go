package netroutine

import (
	"encoding/json"
	"fmt"
)

const idBlockSetVariable = "BlockSetVariable"

type BlockSetVariable struct {
	ToKey string
	Value interface{}
}

func (b *BlockSetVariable) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockSetVariable) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockSetVariable) kind() string {
	return idBlockSetVariable
}

func (b *BlockSetVariable) Run(wce *Environment) (string, error) {
	wce.setData(b.ToKey, b.Value)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, b.Value), Success)
}
