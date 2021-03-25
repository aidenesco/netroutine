package netroutine

import (
	"encoding/json"
	"fmt"
	"math"
)

const idBlockMathFloor = "BlockMathFloor"

type BlockMathFloor struct {
	SourceKey string
	ToKey     string
}

func (b *BlockMathFloor) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockMathFloor) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockMathFloor) kind() string {
	return idBlockMathFloor
}

func (b *BlockMathFloor) Run(wce *Environment) (string, error) {

	s, ok := wce.getData(b.SourceKey)
	if !ok {
		return log(b, "couldn't find the source variable", Error)
	}

	sf, err := toFloat64(s)
	if err != nil {
		return log(b, "source variable wasn't a float", Error)
	}

	v := math.Floor(sf)

	wce.setData(b.ToKey, v)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, v), Success)
}
