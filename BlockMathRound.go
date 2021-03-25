package netroutine

import (
	"encoding/json"
	"fmt"
	"math"
)

const IDBlockMathRound = "BlockMathRound"

func init() {
	blocks[IDBlockMathRound] = BlockMathRound{}
}

type BlockMathRound struct {
	FromKey string
	ToKey   string
}

func (b *BlockMathRound) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockMathRound) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockMathRound) kind() string {
	return IDBlockMathRound
}

func (b *BlockMathRound) Run(wce *Environment) (string, error) {

	s, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "couldn't find the source variable", Error)
	}

	sf, err := toFloat64(s)
	if err != nil {
		return log(b, "source variable wasn't a float", Error)
	}

	v := math.Round(sf)

	wce.setData(b.ToKey, v)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, v), Success)
}
