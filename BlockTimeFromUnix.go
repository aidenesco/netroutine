package netroutine

import (
	"encoding/json"
	"fmt"
	"time"
)

const idBlockTimeFromUnix = "BlockTimeFromUnix"

type BlockTimeFromUnix struct {
	ToKey   string
	FromKey string
}

func (b *BlockTimeFromUnix) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockTimeFromUnix) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockTimeFromUnix) kind() string {
	return idBlockTimeFromUnix
}

func (b *BlockTimeFromUnix) Run(wce *Environment) (string, Status) {
	val, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "variable not found", Error)
	}

	tUnix, err := toInt64(val)
	if err != nil {
		return log(b, fmt.Sprintf("error converting to Int64: %v", err), Error)
	}

	uTime := time.Unix(tUnix, 0)

	wce.setData(b.ToKey, uTime)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, uTime.String()), Success)
}
