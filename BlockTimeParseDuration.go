package netroutine

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const idBlockTimeParseDuration = "BlockTimeParseDuration"

type BlockTimeParseDuration struct {
	ToKey   string
	FromKey string
}

func (b *BlockTimeParseDuration) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockTimeParseDuration) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockTimeParseDuration) kind() string {
	return idBlockTimeParseDuration
}

func (b *BlockTimeParseDuration) Run(wce *Environment) (string, Status) {
	from, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "unable to find variable to parse", Error)
	}
	var dur = time.Second

	switch f := from.(type) {
	case string:
		p, err := strconv.ParseInt(f, 10, 64)
		if err != nil {
			return log(b, fmt.Sprintf("error parsing string to int: %v", err), Error)
		}

		dur = dur * time.Duration(p)
	case int64:
		dur = dur * time.Duration(f)
	case float64:
		dur = dur * time.Duration(int64(f))
	default:
		return log(b, "unable type to convert", Error)
	}

	wce.setData(b.ToKey, dur)

	return log(b, fmt.Sprintf("set %v to a duration of %v seconds", b.ToKey, dur.Seconds()), Success)
}
