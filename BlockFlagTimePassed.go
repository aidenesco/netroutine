package netroutine

import (
	"encoding/json"
	"fmt"
	"time"
)

const idBlockFlagTimePassed = "BlockFlagTimePassed"

type BlockFlagTimePassed struct {
	FromKey string
	ToKey   string
}

func (b *BlockFlagTimePassed) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockFlagTimePassed) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockFlagTimePassed) kind() string {
	return idBlockFlagTimePassed
}

func (b *BlockFlagTimePassed) Run(wce *Environment) (string, error) {
	current := time.Now()

	val, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "unable to retrieve value", Error)
	}

	tval, err := toTime(val)
	if err != nil {
		return log(b, "unable to convert value to Time", Error)
	}

	passed := current.Before(tval)

	wce.setExportData(b.ToKey, passed)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, passed), Success)
}
