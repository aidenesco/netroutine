package netroutine

import (
	"encoding/json"
	"fmt"
	"time"
)

const idBlockTimeToUnix = "BlockTimeToUnix"

type BlockTimeToUnix struct {
	ToKey   string
	FromKey string
}

func (b *BlockTimeToUnix) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockTimeToUnix) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockTimeToUnix) kind() string {
	return idBlockTimeToUnix
}

func (b *BlockTimeToUnix) Run(wce *Environment) (string, Status) {
	oTime, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "variable not found", Error)
	}

	nTime, ok := oTime.(time.Time)
	if !ok {
		return log(b, "variable not time.Time", Error)
	}

	wce.setData(b.ToKey, nTime.Unix())

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, nTime.Unix()), Success)
}
