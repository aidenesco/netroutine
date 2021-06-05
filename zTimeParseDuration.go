package netroutine

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
)

const idTimeParseDuration = "TimeParseDuration"

func init() {
	blocks[idTimeParseDuration] = &TimeParseDuration{}
}

type TimeParseDuration struct {
	ToKey   string
	FromKey string
}

func (b *TimeParseDuration) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *TimeParseDuration) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *TimeParseDuration) kind() string {
	return idTimeParseDuration
}

func (b *TimeParseDuration) Run(ctx context.Context, wce *Environment) (string, Status) {
	from, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}
	var dur = time.Second

	switch f := from.(type) {
	case string:
		p, err := strconv.ParseInt(f, 10, 64)
		if err != nil {
			return log(b, reportError("string to int", err), Error)
		}

		dur = dur * time.Duration(p)
	case int64:
		dur = dur * time.Duration(f)
	case float64:
		dur = dur * time.Duration(int64(f))
	default:
		return log(b, reportWrongType(b.FromKey), Error)
	}

	wce.setData(b.ToKey, dur)

	return log(b, setWorkingData(b.ToKey, dur.String()), Success)
}
