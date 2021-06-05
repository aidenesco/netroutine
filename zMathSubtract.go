package netroutine

import (
	"context"
	"encoding/json"
)

const idMathSubtract = "MathSubtract"

func init() {
	blocks[idMathSubtract] = &MathSubtract{}
}

type MathSubtract struct {
	Source1Key string
	Source2Key string
	ToKey      string
}

func (b *MathSubtract) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *MathSubtract) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *MathSubtract) kind() string {
	return idMathSubtract
}

func (b *MathSubtract) Run(ctx context.Context, wce *Environment) (string, Status) {
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

	v := s1f - s2f

	wce.setData(b.ToKey, v)

	return log(b, setWorkingData(b.ToKey, v), Success)
}
