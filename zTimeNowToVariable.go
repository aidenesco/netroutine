package netroutine

import (
	"context"
	"encoding/json"
	"time"
)

const (
	idTimeNowToVariable = "TimeNowToVariable"
)

func init() {
	blocks[idTimeNowToVariable] = &TimeNowToVariable{}
}

type TimeNowToVariable struct {
	ToKey string
}

func (b *TimeNowToVariable) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *TimeNowToVariable) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *TimeNowToVariable) kind() string {
	return idTimeNowToVariable
}

func (b *TimeNowToVariable) Run(ctx context.Context, wce *Environment) (string, Status) {
	wce.setData(b.ToKey, time.Now())

	return log(b, setWorkingData(b.ToKey, "now"), Success)
}
