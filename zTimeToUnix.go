package netroutine

import (
	"context"
	"encoding/json"
	"time"
)

const idTimeToUnix = "TimeToUnix"

func init() {
	blocks[idTimeToUnix] = &TimeToUnix{}
}

type TimeToUnix struct {
	ToKey   string
	FromKey string
}

func (b *TimeToUnix) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *TimeToUnix) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *TimeToUnix) kind() string {
	return idTimeToUnix
}

func (b *TimeToUnix) Run(ctx context.Context, wce *Environment) (string, Status) {
	oTime, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	nTime, ok := oTime.(time.Time)
	if !ok {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	wce.setData(b.ToKey, nTime.Unix())

	return log(b, setWorkingData(b.ToKey, nTime.String()), Success)
}
