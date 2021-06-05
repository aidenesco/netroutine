package netroutine

import (
	"context"
	"encoding/json"
)

const idMathTotal = "MathTotal"

func init() {
	blocks[idMathTotal] = &MathTotal{}
}

type MathTotal struct {
	FromKey string
	ToKey   string
}

func (b *MathTotal) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *MathTotal) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *MathTotal) kind() string {
	return idMathTotal
}

func (b *MathTotal) Run(ctx context.Context, wce *Environment) (string, Status) {
	var total float64
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	s, ok := v.([]interface{})
	if !ok {
		f, err := toFloat64(v)
		if err != nil {
			return log(b, reportWrongType(b.FromKey), Error)
		}
		total = f

		wce.setData(b.ToKey, total)

		return log(b, setWorkingData(b.ToKey, total), Success)
	}

	for _, v := range s {
		f, err := toFloat64(v)
		if err != nil {
			continue
		}
		total += f
	}

	wce.setData(b.ToKey, total)

	return log(b, setWorkingData(b.ToKey, total), Success)
}
