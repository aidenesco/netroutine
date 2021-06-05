package netroutine

import (
	"context"
	"encoding/json"
	"math"
)

const idMathFloor = "MathFloor"

func init() {
	blocks[idMathFloor] = &MathFloor{}
}

type MathFloor struct {
	FromKey string
	ToKey   string
}

func (b *MathFloor) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *MathFloor) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *MathFloor) kind() string {
	return idMathFloor
}

func (b *MathFloor) Run(ctx context.Context, wce *Environment) (string, Status) {

	s, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	sf, err := toFloat64(s)
	if err != nil {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	v := math.Floor(sf)

	wce.setData(b.ToKey, v)

	return log(b, setWorkingData(b.ToKey, v), Success)
}
