package netroutine

import (
	"encoding/json"
	"fmt"
	"time"
)

const idBlockTimeFlagPassed = "BlockTimeFlagPassed"

type BlockTimeFlagPassed struct {
	FromKey string
	ToKey   string
}

func (b *BlockTimeFlagPassed) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockTimeFlagPassed) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockTimeFlagPassed) kind() string {
	return idBlockTimeFlagPassed
}

func (b *BlockTimeFlagPassed) Run(wce *Environment) (string, Status) {
	val, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "unable to retrieve value", Error)
	}

	oTime, ok := val.(time.Time)
	if !ok {
		return log(b, "variable not time.Time", Error)
	}

	passed := time.Now().After(oTime)

	wce.setData(b.ToKey, passed)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, passed), Success)
}
