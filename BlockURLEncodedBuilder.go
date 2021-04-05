package netroutine

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const idBlockURLEncodedBuilder = "BlockURLEncodedBuilder"

type BlockURLEncodedBuilder struct {
	Values []URLEncodedValueBuilder
	ToKey  string
}

type URLEncodedValueBuilder struct {
	ToPath    string
	Variables []string
	Value     string
	Complex   bool
}

func (b *BlockURLEncodedBuilder) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockURLEncodedBuilder) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockURLEncodedBuilder) kind() string {
	return idBlockURLEncodedBuilder
}

func (b *BlockURLEncodedBuilder) Run(wce *Environment) (string, Status) {
	data := url.Values{}

	for _, v := range b.Values {
		if !v.Complex {
			data.Set(v.ToPath, v.Value)
			continue
		}

		var sub []interface{}
		for _, va := range v.Variables {
			sv, ok := wce.getData(va)
			if !ok {
				return log(b, fmt.Sprintf("Failed to find \"%v\" variable", va), Error)
			}
			sub = append(sub, sv)
		}

		data.Set(v.ToPath, fmt.Sprintf(v.Value, sub...))
	}

	built := data.Encode()

	wce.setData(b.ToKey, built)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, built), Success)
}
