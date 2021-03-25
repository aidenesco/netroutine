package netroutine

import (
	"encoding/json"
	"fmt"
)

const idBlockMathTotal = "BlockMathTotal"

type BlockMathTotal struct {
	FromKey string
	ToKey   string
}

func (b *BlockMathTotal) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockMathTotal) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockMathTotal) kind() string {
	return idBlockMathTotal
}

func (b *BlockMathTotal) Run(wce *Environment) (string, error) {
	var total float64
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "couldn't find the variable", Error)
	}

	s, ok := v.([]interface{})
	if !ok {
		f, err := toFloat64(v)
		if err != nil {
			return log(b, "variable not a slice or a float", Error)
		}
		total = f

		wce.setData(b.ToKey, total)

		return log(b, fmt.Sprintf("set %v to %v", b.ToKey, total), Success)
	}

	for _, v := range s {
		f, err := toFloat64(v)
		if err != nil {
			continue
		}
		total += f
	}

	wce.setData(b.ToKey, total)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, total), Success)
}
