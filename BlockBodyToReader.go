package netroutine

import (
	"encoding/json"
	"fmt"
	"strings"
)

const idBlockBodyToReader = "BodyToReader"

type BlockBodyToReader struct {
	ToKey string
}

func (b *BlockBodyToReader) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockBodyToReader) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockBodyToReader) kind() string {
	return idBlockBodyToReader
}

func (b *BlockBodyToReader) Run(wce *Environment) (string, Status) {
	body, err := wce.lastResponseBody()
	if err != nil {
		return log(b, "couldn't get last response body", Error)
	}

	wce.setData(b.ToKey, strings.NewReader(body))

	return log(b, fmt.Sprintf("set %v to a reader", b.ToKey), Success)
}
