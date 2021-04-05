package netroutine

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const idBlockURLDecode = "BlockURLDecode"

type BlockURLDecode struct {
	FromKey string
	ToKey   string
}

func (b *BlockURLDecode) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockURLDecode) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockURLDecode) kind() string {
	return idBlockURLDecode
}

func (b *BlockURLDecode) Run(wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "couldn't find the variable to decode", Error)
	}

	s, err := toString(v)
	if err != nil {
		return log(b, "unable to convert variable to string", Error)
	}

	decodedValue, err := url.QueryUnescape(s)
	if err != nil {
		return log(b, fmt.Sprintf("error decoding string - %v", err), Error)
	}

	wce.setData(b.ToKey, decodedValue)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, decodedValue), Success)
}
