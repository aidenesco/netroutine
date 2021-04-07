package netroutine

import (
	"encoding/json"
	"fmt"
	"time"
)

const idBlockTimeParseFormatted = "BlockTimeParseFormatted"

type BlockTimeParseFormatted struct {
	FromKey string
	Format  string
	ToKey   string
}

func (b *BlockTimeParseFormatted) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockTimeParseFormatted) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockTimeParseFormatted) kind() string {
	return idBlockTimeParseFormatted
}

func (b *BlockTimeParseFormatted) Run(wce *Environment) (string, Status) {
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

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, ptime.String()), Success)
}
