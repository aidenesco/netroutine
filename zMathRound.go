package netroutine

import (
	"context"
	"encoding/json"
	"math"
)

const idMathRound = "MathRound"

func init() {
	blocks[idMathRound] = &MathRound{}
}

type MathRound struct {
	FromKey string
	ToKey   string
}

func (b *MathRound) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *MathRound) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *MathRound) kind() string {
	return idMathRound
}

func (b *MathRound) Run(ctx context.Context, wce *Environment) (string, Status) {
	s, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	sf, err := toFloat64(s)
	if err != nil {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	v := math.Round(sf)

	wce.setData(b.ToKey, v)

	return log(b, setWorkingData(b.ToKey, v), Success)
}
