package netroutine

import (
	"context"
	"encoding/json"
	"time"
)

const idTimeAddDuration = "TimeAddDuration"

func init() {
	blocks[idTimeAddDuration] = &TimeAddDuration{}
}

type TimeAddDuration struct {
	ToKey       string
	TimeKey     string
	DurationKey string
}

func (b *TimeAddDuration) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *TimeAddDuration) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *TimeAddDuration) kind() string {
	return idTimeAddDuration
}

func (b *TimeAddDuration) Run(ctx context.Context, wce *Environment) (string, Status) {
	t, ok := wce.getData(b.TimeKey)
	if !ok {
		return log(b, missingWorkingData(b.TimeKey), Error)
	}

	oTime, err := toTime(t)
	if err != nil {
		return log(b, reportError("to time", err), Error)
	}

	d, ok := wce.getData(b.DurationKey)
	if !ok {
		return log(b, missingWorkingData(b.DurationKey), Error)
	}

	oDuration, ok := d.(time.Duration)
	if !ok {
		return log(b, reportWrongType(b.DurationKey), Error)
	}

	nTime := oTime.Add(oDuration)

	wce.setData(b.ToKey, nTime)

	return log(b, setWorkingData(b.ToKey, nTime.String()), Success)
}
