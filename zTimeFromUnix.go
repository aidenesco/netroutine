package netroutine

import (
	"context"
	"encoding/json"
	"time"
)

const (
	idTimeFromUnix = "TimeFromUnix"
)

func init() {
	blocks[idTimeFromUnix] = &TimeFromUnix{}
}

type TimeFromUnix struct {
	ToKey   string
	FromKey string
}

func (b *TimeFromUnix) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *TimeFromUnix) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *TimeFromUnix) kind() string {
	return idTimeFromUnix
}

func (b *TimeFromUnix) Run(ctx context.Context, wce *Environment) (string, Status) {
	val, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	tUnix, err := toInt64(val)
	if err != nil {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	uTime := time.Unix(tUnix, 0)

	wce.setData(b.ToKey, uTime)

	return log(b, setWorkingData(b.ToKey, uTime.String()), Success)
}
