package netroutine

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

const (
	idGenerateAppId = "GenerateAppId"
)

func init() {
	blocks[idGenerateAppId] = &GenerateAppId{}
}

type GenerateAppId struct {
	ToKey string
}

func (b *GenerateAppId) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *GenerateAppId) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *GenerateAppId) kind() string {
	return idGenerateAppId
}

func (b *GenerateAppId) Run(ctx context.Context, wce *Environment) (string, Status) {
	s := strings.ToUpper(uuid.New().String())

	wce.setData(b.ToKey, s)

	return log(b, setWorkingData(b.ToKey, s), Success)
}
