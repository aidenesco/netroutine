package netroutine

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const idBlockUnixMilli = "BlockUnixMilli"

type BlockUnixMilli struct {
	ToKey string
}

func (b *BlockUnixMilli) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockUnixMilli) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockUnixMilli) kind() string {
	return idBlockUnixMilli
}

func (b *BlockUnixMilli) Run(wce *Environment) (string, error) {
	ts := strconv.FormatInt(time.Now().UTC().UnixNano()/1e6, 10)

	wce.setData(b.ToKey, ts)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, ts), Success)
}
