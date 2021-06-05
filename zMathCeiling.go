package netroutine

import (
	"context"
	"encoding/json"
	"math"
)

const idMathCeiling = "MathCeiling"

func init() {
	blocks[idMathCeiling] = &MathCeiling{}
}

type MathCeiling struct {
	FromKey string
	ToKey   string
}

func (b *MathCeiling) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *MathCeiling) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *MathCeiling) kind() string {
	return idMathCeiling
}

func (b *MathCeiling) Run(ctx context.Context, wce *Environment) (string, Status) {

	s, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	sf, err := toFloat64(s)
	if err != nil {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	v := math.Ceil(sf)

	wce.setData(b.ToKey, v)

	return log(b, setWorkingData(b.ToKey, v), Success)
}
