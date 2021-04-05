package netroutine

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const idBlockUnixNano = "BlockUnixNano"

type BlockUnixNano struct {
	ToKey string
}

func (b *BlockUnixNano) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockUnixNano) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockUnixNano) kind() string {
	return idBlockUnixNano
}

func (b *BlockUnixNano) Run(wce *Environment) (string, Status) {
	ts := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	wce.setData(b.ToKey, ts)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, ts), Success)
}
