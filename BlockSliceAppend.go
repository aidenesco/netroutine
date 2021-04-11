package netroutine

import (
	"encoding/json"
	"fmt"
)

func init() {
	blocks[idBlockSliceAppend] = BlockSliceAppend{}
}

const idBlockSliceAppend = "BlockSliceAppend"

type BlockSliceAppend struct {
	ToKey   string
	FromKey string
}

func (b *BlockSliceAppend) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockSliceAppend) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockSliceAppend) kind() string {
	return idBlockSliceAppend
}

func (b *BlockSliceAppend) Run(wce *Environment) (string, Status) {
	v, ok := wce.getData(b.ToKey)
	if !ok {
		var newS []interface{}
		wce.setData(b.ToKey, newS)
		v = newS
	}

	sl, ok := v.([]interface{})
	if !ok {
		return log(b, "variable not a slice", Error)
	}

	v, ok = wce.getData(b.FromKey)
	if !ok {
		return log(b, "unable to find source variable", Error)
	}

	sl = append(sl, v)

	wce.setData(b.ToKey, sl)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, sl), Success)
}
