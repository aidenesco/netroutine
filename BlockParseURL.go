package netroutine

import (
	"encoding/json"
	"fmt"
)

const idBlockParseURL = "BlockParseURL"

type BlockParseURL struct {
	ToKey string
}

func (b *BlockParseURL) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockParseURL) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockParseURL) kind() string {
	return idBlockParseURL
}

func (b *BlockParseURL) Run(wce *Environment) (string, Status) {
	resp, err := wce.lastResponse()
	if err != nil {
		return log(b, fmt.Sprintf("error getting response - %v", err), Error)
	}

	urls := resp.Request.URL.String()

	wce.setData(b.ToKey, urls)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, urls), Success)
}
