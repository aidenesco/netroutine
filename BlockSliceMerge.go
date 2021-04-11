package netroutine

import (
	"encoding/json"
	"fmt"
)

const idBlockSliceMerge = "BlockSliceMerge"

type BlockSliceMerge struct {
	ToKey    string
	Format   string
	FromKeys []string
}

func (b *BlockSliceMerge) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockSliceMerge) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockSliceMerge) kind() string {
	return idBlockSliceMerge
}

func (b *BlockSliceMerge) Run(wce *Environment) (string, Status) {
	var built []string
	var sources [][]interface{}
	var length int

	for _, v := range b.FromKeys {
		data, ok := wce.getData(v)
		if !ok {
			return log(b, "unable to find source variable", Error)
		}

		sl, ok := data.([]interface{})
		if !ok {
			return log(b, "variable not a slice", Error)
		}

		sources = append(sources, sl)
	}

	if len(sources) == 0 {
		return log(b, "no sources found", Error)
	}

	length = len(sources[0])

	for i := 1; i < length; i++ {
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

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, built), Success)
}
