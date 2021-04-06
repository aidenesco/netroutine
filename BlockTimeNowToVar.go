package netroutine

import (
	"encoding/json"
	"fmt"
	"time"
)

const idBlockTimeNowToVar = "TimeNowToVar"

type BlockTimeNowToVar struct {
	ToKey string
}

func (b *BlockTimeNowToVar) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockTimeNowToVar) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockTimeNowToVar) kind() string {
	return idBlockTimeNowToVar
}

func (b *BlockTimeNowToVar) Run(wce *Environment) (string, Status) {
	wce.setData(b.ToKey, time.Now())

	return log(b, fmt.Sprintf("set %v to current time", b.ToKey), Success)
}
