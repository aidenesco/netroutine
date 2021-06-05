package netroutine

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

const idBuilderURLEncoded = "BuilderURLEncoded"

func init() {
	blocks[idBuilderURLEncoded] = &BuilderURLEncoded{}
}

type BuilderURLEncoded struct {
	Values []struct {
		ToPath    string
		Variables []string
		Value     string
		Complex   bool
	}
	ToKey string
}

func (b *BuilderURLEncoded) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BuilderURLEncoded) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BuilderURLEncoded) kind() string {
	return idBuilderURLEncoded
}

func (b *BuilderURLEncoded) Run(ctx context.Context, wce *Environment) (string, Status) {
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
				return log(b, missingWorkingData(va), Error)
			}
			sub = append(sub, sv)
		}

		data.Set(v.ToPath, fmt.Sprintf(v.Value, sub...))
	}

	built := data.Encode()

	wce.setData(b.ToKey, built)

	return log(b, setWorkingData(b.ToKey, built), Success)
}
