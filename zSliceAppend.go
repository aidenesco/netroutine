package netroutine

import (
	"context"
	"encoding/json"
	"fmt"
)

func init() {
	blocks[idSliceAppend] = &SliceAppend{}
}

const idSliceAppend = "SliceAppend"

type SliceAppend struct {
	ToKey   string
	FromKey string
}

func (b *SliceAppend) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *SliceAppend) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *SliceAppend) kind() string {
	return idSliceAppend
}

func (b *SliceAppend) Run(ctx context.Context, wce *Environment) (string, Status) {
	v, ok := wce.getData(b.ToKey)
	if !ok {
		var newS []interface{}
		wce.setData(b.ToKey, newS)
		v = newS
	}

	sl, ok := v.([]interface{})
	if !ok {
		return log(b, reportWrongType(b.ToKey), Error)
	}

	v, ok = wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	sl = append(sl, v)

	wce.setData(b.ToKey, sl)

	return log(b, setWorkingData(b.ToKey, fmt.Sprintf("%v", sl)), Success)
}
