package netroutine

import (
	"encoding/json"
	"fmt"
)

const idBlockStringBuilder = "BlockStringBuilder"

type BlockStringBuilder struct {
	Variables []string
	Base      string
	ToKey     string
}

func (b *BlockStringBuilder) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockStringBuilder) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockStringBuilder) kind() string {
	return idBlockStringBuilder
}

func (b *BlockStringBuilder) Run(wce *Environment) (string, error) {
	var sub []interface{}
	for _, v := range b.Variables {
		sv, ok := wce.getData(v)
		if !ok {
			return log(b, fmt.Sprintf("Failed to find \"%v\" variable", v), Error)
		}
		sub = append(sub, sv)
	}

	fs := fmt.Sprintf(b.Base, sub...)

	wce.setData(b.ToKey, fs)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, fs), Success)
}
