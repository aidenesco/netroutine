package netroutine

import (
	"encoding/json"
	"fmt"
)

const idBlockMathMultiply = "BlockMathMultiply"

type BlockMathMultiply struct {
	Source1Key string
	Source2Key string
	ToKey      string
}

func (b *BlockMathMultiply) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockMathMultiply) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockMathMultiply) kind() string {
	return idBlockMathMultiply
}

func (b *BlockMathMultiply) Run(wce *Environment) (string, error) {

	s1, ok := wce.getData(b.Source1Key)
	if !ok {
		return log(b, "couldn't find the first source variable", Error)
	}

	s1f, err := toFloat64(s1)
	if err != nil {
		return log(b, "first source variable wasn't a float", Error)
	}

	s2, ok := wce.getData(b.Source2Key)
	if !ok {
		return log(b, "couldn't find the second source variable", Error)
	}

	s2f, err := toFloat64(s2)
	if err != nil {
		return log(b, "second source variable wasn't a float", Error)
	}

	v := s1f * s2f

	wce.setData(b.ToKey, v)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, v), Success)
}
