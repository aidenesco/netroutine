package netroutine

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const idBlockUnix = "BlockUnix"

type BlockUnix struct {
	ToKey string
}

func (b *BlockUnix) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockUnix) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockUnix) kind() string {
	return idBlockUnix
}

func (b *BlockUnix) Run(wce *Environment) (string, Status) {
	ts := strconv.FormatInt(time.Now().UTC().Unix(), 10)

	wce.setData(b.ToKey, ts)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, ts), Success)
}
