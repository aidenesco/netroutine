package netroutine

import (
	"context"
	"encoding/json"
	"net/url"
)

const idEncodeURL = "EncodeURL"

func init() {
	blocks[idEncodeURL] = &EncodeURL{}
}

type EncodeURL struct {
	FromKey string
	ToKey   string
}

func (b *EncodeURL) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *EncodeURL) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *EncodeURL) kind() string {
	return idEncodeURL
}

func (b *EncodeURL) Run(ctx context.Context, wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	s, err := toString(v)
	if err != nil {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	enc := url.QueryEscape(s)

	wce.setData(b.ToKey, enc)

	return log(b, setWorkingData(b.ToKey, enc), Success)
}
