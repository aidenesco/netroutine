package netroutine

import (
	"context"
	"encoding/json"

	"github.com/Jeffail/gabs/v2"
)

const idBuilderJSON = "BuilderJSON"

func init() {
	blocks[idBuilderJSON] = &BuilderJSON{}
}

type BuilderJSON struct {
	Values []struct {
		ToPath   []string
		Variable string
		Value    interface{}
		Complex  bool
	}
	ToKey string
}

func (b *BuilderJSON) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BuilderJSON) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BuilderJSON) kind() string {
	return idBuilderJSON
}

func (b *BuilderJSON) Run(ctx context.Context, wce *Environment) (string, Status) {
	data := gabs.New()

	for _, v := range b.Values {
		if !v.Complex {
			_, err := data.Set(v.Value, v.ToPath...)
			if err != nil {
				return log(b, reportError("setting json", err), Error)
			}
			continue
		}

		val, ok := wce.getData(v.Variable)
		if !ok {
			return log(b, missingWorkingData(v.Variable), Error)
		}

		_, err := data.Set(val, v.ToPath...)
		if err != nil {
			return log(b, reportError("setting json", err), Error)
		}
	}

	built := data.String()

	wce.setData(b.ToKey, built)

	return log(b, setWorkingData(b.ToKey, built), Success)
}
