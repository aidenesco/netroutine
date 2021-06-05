package netroutine

import (
	"context"
	"encoding/json"
)

const idMathMultiply = "MathMultiply"

func init() {
	blocks[idMathMultiply] = &MathMultiply{}
}

type MathMultiply struct {
	Source1Key string
	Source2Key string
	ToKey      string
}

func (b *MathMultiply) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *MathMultiply) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *MathMultiply) kind() string {
	return idMathMultiply
}

func (b *MathMultiply) Run(ctx context.Context, wce *Environment) (string, Status) {

	s1, ok := wce.getData(b.Source1Key)
	if !ok {
		return log(b, missingWorkingData(b.Source1Key), Error)
	}

	s1f, err := toFloat64(s1)
	if err != nil {
		return log(b, reportWrongType(b.Source1Key), Error)
	}

	s2, ok := wce.getData(b.Source2Key)
	if !ok {
		return log(b, missingWorkingData(b.Source2Key), Error)
	}

	s2f, err := toFloat64(s2)
	if err != nil {
		return log(b, reportWrongType(b.Source2Key), Error)
	}

	v := s1f * s2f

	wce.setData(b.ToKey, v)

	return log(b, setWorkingData(b.ToKey, v), Success)
}
