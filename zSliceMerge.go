package netroutine

import (
	"context"
	"encoding/json"
	"fmt"
)

func init() {
	blocks[idSliceMerge] = &SliceMerge{}
}

const idSliceMerge = "SliceMerge"

type SliceMerge struct {
	ToKey    string
	Format   string
	FromKeys []string
}

func (b *SliceMerge) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *SliceMerge) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *SliceMerge) kind() string {
	return idSliceMerge
}

func (b *SliceMerge) Run(ctx context.Context, wce *Environment) (string, Status) {
	built := []string{}
	var sources [][]interface{}
	var length int

	for _, v := range b.FromKeys {
		data, ok := wce.getData(v)
		if !ok {
			return log(b, missingWorkingData(v), Error)
		}

		sl, ok := data.([]interface{})
		if !ok {
			return log(b, reportWrongType(v), Error)
		}

		sources = append(sources, sl)
	}

	if len(sources) == 0 {
		return log(b, "no sources found", Error)
	}

	length = len(sources[0])

	for i := 1; i < len(sources); i++ {
		if len(sources[i]) != length {
			return log(b, "got slices of varying lengths", Error)
		}
	}

	for i := 0; i < length; i++ {
		var vars []interface{}
		for _, v := range sources {
			vars = append(vars, v[i])
		}

		built = append(built, fmt.Sprintf(b.Format, vars...))
	}

	wce.setData(b.ToKey, built)

	return log(b, setWorkingData(b.ToKey, fmt.Sprintf("%v", built)), Success)
}
