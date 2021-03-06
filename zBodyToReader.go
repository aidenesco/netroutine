package netroutine

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

const idBodyToReader = "BodyToReader"

func init() {
	blocks[idBodyToReader] = &BodyToReader{}
}

type BodyToReader struct {
	ToKey string
}

func (b *BodyToReader) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BodyToReader) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BodyToReader) kind() string {
	return idBodyToReader
}

func (b *BodyToReader) Run(ctx context.Context, wce *Environment) (string, Status) {
	if wce.lastResponseBody == "" {
		return log(b, "getting response body", Error)
	}

	wce.setData(b.ToKey, strings.NewReader(wce.lastResponseBody))

	return log(b, fmt.Sprintf("set \"%v\" to a reader", b.ToKey), Success)
}
