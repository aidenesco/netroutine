package netroutine

import (
	"context"
	"encoding/json"
)

const idMathAdd = "MathAdd"

func init() {
	blocks[idMathAdd] = &MathAdd{}
}

type MathAdd struct {
	Source1Key string
	Source2Key string
	ToKey      string
}

func (b *MathAdd) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *MathAdd) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *MathAdd) kind() string {
	return idMathAdd
}

func (b *MathAdd) Run(ctx context.Context, wce *Environment) (string, Status) {

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

	v := s1f + s2f

	wce.setData(b.ToKey, v)

	return log(b, setWorkingData(b.ToKey, v), Success)
}
