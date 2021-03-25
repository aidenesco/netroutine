package netroutine

import (
	"encoding/json"
	"fmt"
)

const idBlockParseHeader = "BlockParseHeader"

type BlockParseHeader struct {
	Header   string
	ToKey    string
	Required bool
}

func (b *BlockParseHeader) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockParseHeader) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockParseHeader) kind() string {
	return idBlockParseHeader
}

func (b *BlockParseHeader) Run(wce *Environment) (string, error) {
	resp, err := wce.lastResponse()
	if err != nil {
		return log(b, fmt.Sprintf("error getting response - %v", err), Error)
	}

	val := resp.Header.Get(b.Header)
	if val == "" {
		if b.Required {
			return log(b, "unable to find required header", Fail)
		}
		return log(b, "unable to find non required header", Success)
	}

	wce.setData(b.ToKey, val)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, val), Success)
}
