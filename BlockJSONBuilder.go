package netroutine

import (
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs/v2"
)

const idBlockJSONBuilder = "BlockJSONBuilder"

type BlockJSONBuilder struct {
	Values []JSONValueBuilder
	ToKey  string
}

type JSONValueBuilder struct {
	ToPath   []string
	Variable string
	Value    interface{}
	Complex  bool
}

func (b *BlockJSONBuilder) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockJSONBuilder) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockJSONBuilder) kind() string {
	return idBlockJSONBuilder
}

func (b *BlockJSONBuilder) Run(wce *Environment) (string, Status) {
	data := gabs.New()

	for _, v := range b.Values {
		if !v.Complex {
			_, err := data.Set(v.Value, v.ToPath...)
			if err != nil {
				return log(b, fmt.Sprintf("error setting json: key %v - value %v", v.ToPath, v.Value), Error)
			}
			continue
		}

		val, ok := wce.getData(v.Variable)
		if !ok {
			return log(b, fmt.Sprintf("Failed to find \"%s\" variable", v.Variable), Error)
		}

		_, err := data.Set(val, v.ToPath...)
		if err != nil {
			return log(b, fmt.Sprintf("error setting json: key %v - value %v", v.ToPath, v.Value), Error)
		}
	}

	built := data.String()

	wce.setData(b.ToKey, built)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, built), Success)
}
