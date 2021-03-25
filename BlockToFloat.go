package netroutine

import (
	"encoding/json"
	"fmt"
)

const idBlockToFloat = "BlockToFloat"

type BlockToFloat struct {
	FromKey string
	ToKey   string
}

func (b *BlockToFloat) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockToFloat) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockToFloat) kind() string {
	return idBlockToFloat
}

func (b *BlockToFloat) Run(wce *Environment) (string, error) {
	orig, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "couldn't find the source variable", Error)
	}

	f, err := toFloat64(orig)
	if err != nil {
		return log(b, "unable to parse float", Error)
	}

	wce.setData(b.ToKey, f)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, f), Success)

}
