package netroutine

import (
	"encoding/json"
	"fmt"
	"time"
)

const idBlockTimeAddDuration = "BlockTimeAddDuration"

type BlockTimeAddDuration struct {
	ToKey       string
	TimeKey     string
	DurationKey string
}

func (b *BlockTimeAddDuration) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockTimeAddDuration) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockTimeAddDuration) kind() string {
	return idBlockTimeAddDuration
}

func (b *BlockTimeAddDuration) Run(wce *Environment) (string, Status) {
	t, ok := wce.getData(b.TimeKey)
	if !ok {
		return log(b, "unable to find time variable", Error)
	}

	oTime, ok := t.(time.Time)
	if !ok {
		return log(b, "value not time.Time", Error)
	}

	d, ok := wce.getData(b.DurationKey)
	if !ok {
		return log(b, "unable to find duration variable", Error)
	}

	oDuration, ok := d.(time.Duration)
	if !ok {
		return log(b, "value not time.Duration", Error)
	}

	nTime := oTime.Add(oDuration)

	wce.setData(b.ToKey, nTime)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, nTime.String()), Success)
}
