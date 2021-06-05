package netroutine

import (
	"context"
	"encoding/json"
	"time"
)

const idTimeFromFormatted = "TimeFromFormatted"

func init() {
	blocks[idTimeFromFormatted] = &TimeFromFormatted{}
}

type TimeFromFormatted struct {
	FromKey string
	Format  string
	ToKey   string
}

func (b *TimeFromFormatted) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *TimeFromFormatted) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *TimeFromFormatted) kind() string {
	return idTimeFromFormatted
}

func (b *TimeFromFormatted) Run(ctx context.Context, wce *Environment) (string, Status) {
	old, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	s, err := toString(old)
	if err != nil {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	ptime, err := time.Parse(b.Format, s)
	if err != nil {
		return log(b, reportError("parsing time", err), Error)
	}

	wce.setData(b.ToKey, ptime)

	return log(b, setWorkingData(b.ToKey, ptime.String()), Success)
}
