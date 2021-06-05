package netroutine

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const idTimeFlagPassed = "TimeFlagPassed"

func init() {
	blocks[idTimeFlagPassed] = &TimeFlagPassed{}
}

type TimeFlagPassed struct {
	FromKey string
	ToKey   string
}

func (b *TimeFlagPassed) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *TimeFlagPassed) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *TimeFlagPassed) kind() string {
	return idTimeFlagPassed
}

func (b *TimeFlagPassed) Run(ctx context.Context, wce *Environment) (string, Status) {
	val, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	oTime, ok := val.(time.Time)
	if !ok {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	passed := time.Now().After(oTime)

	wce.setData(b.ToKey, passed)

	return log(b, setWorkingData(b.ToKey, fmt.Sprintf("%v", passed)), Success)
}
