package netroutine

import (
	"context"
	"encoding/json"
	"net/url"
)

const (
	idDecodeURL = "DecodeURL"
)

func init() {
	blocks[idDecodeURL] = &DecodeURL{}
}

type DecodeURL struct {
	FromKey string
	ToKey   string
}

func (b *DecodeURL) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *DecodeURL) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *DecodeURL) kind() string {
	return idDecodeURL
}

func (b *DecodeURL) Run(ctx context.Context, wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	s, err := toString(v)
	if err != nil {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	decodedValue, err := url.QueryUnescape(s)
	if err != nil {
		return log(b, reportError("decoding", err), Error)
	}

	wce.setData(b.ToKey, decodedValue)

	return log(b, setWorkingData(b.ToKey, decodedValue), Success)
}
