package netroutine

import (
	"context"
	"encoding/base64"
	"encoding/json"
)

const (
	idDecodeBase64 = "DecodeBase64"
)

func init() {
	blocks[idDecodeBase64] = &DecodeBase64{}
}

type DecodeBase64 struct {
	FromKey string
	ToKey   string
}

func (b *DecodeBase64) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *DecodeBase64) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *DecodeBase64) kind() string {
	return idDecodeBase64
}

func (b *DecodeBase64) Run(ctx context.Context, wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	s, err := toString(v)
	if err != nil {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	e, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return log(b, reportError("decoding", err), Error)
	}

	es := string(e)

	wce.setData(b.ToKey, es)

	return log(b, setWorkingData(b.ToKey, es), Success)
}
