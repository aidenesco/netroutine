package netroutine

import (
	"encoding/json"
	"fmt"
	"time"
)

const idBlockParseTime = "BlockParseTime"

type BlockParseTime struct {
	FromKey string
	Format  string
	ToKey   string
}

func (b *BlockParseTime) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockParseTime) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockParseTime) kind() string {
	return idBlockParseTime
}

func (b *BlockParseTime) Run(wce *Environment) (string, Status) {
	old, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "unable to find value to parse", Error)
	}

	s, err := toString(old)
	if err != nil {
		return log(b, "unable to convert variable to string", Error)
	}

	ptime, err := time.Parse(b.Format, s)
	if err != nil {
		return log(b, fmt.Sprintf("error parsing time - %v", err), Error)
	}

	wce.setData(b.ToKey, ptime)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, ptime), Success)
}
