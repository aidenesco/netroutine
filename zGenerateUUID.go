package netroutine

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

const (
	idGenerateUUID = "GenerateUUID"
)

func init() {
	blocks[idGenerateUUID] = &GenerateUUID{}
}

type GenerateUUID struct {
	ToKey string
}

func (b *GenerateUUID) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *GenerateUUID) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *GenerateUUID) kind() string {
	return idGenerateUUID
}

func (b *GenerateUUID) Run(ctx context.Context, wce *Environment) (string, Status) {
	s := uuid.New().String()

	wce.setData(b.ToKey, s)

	return log(b, setWorkingData(b.ToKey, s), Success)
}
