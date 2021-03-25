package netroutine

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const idBlockURLEncode = "BlockURLEncode"

type BlockURLEncode struct {
	FromKey string
	ToKey   string
}

func (b *BlockURLEncode) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockURLEncode) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockURLEncode) kind() string {
	return idBlockURLEncode
}

func (b *BlockURLEncode) Run(wce *Environment) (string, error) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "couldn't find the variable to encode", Error)
	}

	s, err := toString(v)
	if err != nil {
		return log(b, "unable to convert variable to string", Error)
	}

	enc := url.QueryEscape(s)

	wce.setData(b.ToKey, enc)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, enc), Success)
}
