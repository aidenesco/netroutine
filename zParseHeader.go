package netroutine

import (
	"context"
	"encoding/json"
)

const idParseHeader = "ParseHeader"

func init() {
	blocks[idParseHeader] = &ParseHeader{}
}

type ParseHeader struct {
	Header   string
	ToKey    string
	Required bool
}

func (b *ParseHeader) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *ParseHeader) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *ParseHeader) kind() string {
	return idParseHeader
}

func (b *ParseHeader) Run(ctx context.Context, wce *Environment) (string, Status) {
	if wce.lastResponse == nil {
		return log(b, "getting response", Error)
	}

	val := wce.lastResponse.Header.Get(b.Header)
	if val == "" {
		if b.Required {
			return log(b, "unable to find required header", Fail)
		}
		return log(b, "unable to find non required header", Success)
	}

	wce.setData(b.ToKey, val)

	return log(b, setWorkingData(b.ToKey, val), Success)
}
