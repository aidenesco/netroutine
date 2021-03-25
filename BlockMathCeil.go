package netroutine

import (
	"encoding/json"
	"fmt"
	"math"
)

const idBlockMathCeil = "BlockMathCeil"

type BlockMathCeil struct {
	FromKey string
	ToKey   string
}

func (b *BlockMathCeil) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockMathCeil) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockMathCeil) kind() string {
	return idBlockMathCeil
}

func (b *BlockMathCeil) Run(wce *Environment) (string, error) {

	s, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "couldn't find the source variable", Error)
	}

	sf, err := toFloat64(s)
	if err != nil {
		return log(b, "source variable wasn't a float", Error)
	}

	v := math.Ceil(sf)

	wce.setData(b.ToKey, v)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, v), Success)
}
