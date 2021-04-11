package netroutine

import (
	"encoding/json"
)

func init() {
	blocks[idBlockSliceToSubroutine] = BlockSliceToSubroutine{}
}

const idBlockSliceToSubroutine = "SliceToSubroutine"

type BlockSliceToSubroutine struct {
	VariableKeys []string
	Routine      Routine
	FromKeys     []string
}

func (b *BlockSliceToSubroutine) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockSliceToSubroutine) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockSliceToSubroutine) kind() string {
	return idBlockSliceMerge
}

func (b *BlockSliceToSubroutine) Run(wce *Environment) (string, Status) {
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

	if length != len(b.VariableKeys) {
		return log(b, "got variable keys and sources of different lengths", Error)
	}

	for i := 0; i < length; i++ {
		for j, v := range sources {
			wce.setData(b.VariableKeys[j], v[i])
		}

		b.Routine.Run(wce)

		if wce.Status != Success {
			log(b, "got a status other than Success", wce.Status)
		}
	}

	return log(b, "ran subroutines", Success)
}
