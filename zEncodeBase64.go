package netroutine

import (
	"context"
	"encoding/base64"
	"encoding/json"
)

const idEncodeBase64 = "EncodeBase64"

func init() {
	blocks[idEncodeBase64] = &EncodeBase64{}
}

type EncodeBase64 struct {
	FromKey string
	ToKey   string
}

func (b *EncodeBase64) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *EncodeBase64) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *EncodeBase64) kind() string {
	return idEncodeBase64
}

func (b *EncodeBase64) Run(ctx context.Context, wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	s, err := toString(v)
	if err != nil {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	e := base64.StdEncoding.EncodeToString([]byte(s))

	wce.setData(b.ToKey, e)

	return log(b, setWorkingData(b.ToKey, e), Success)
}
