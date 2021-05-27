package netroutine

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

const (
	idGenerateAppId = "GenerateAppId"
	idGenerateUUID  = "GenerateUUID"
)

func init() {
	blocks[idGenerateAppId] = &GenerateAppId{}
	blocks[idGenerateUUID] = &GenerateUUID{}
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
