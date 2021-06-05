package netroutine

import (
	"context"
	"encoding/json"
	"fmt"
)

const idSetVariable = "SetVariable"

func init() {
	blocks[idSetVariable] = &SetVariable{}
}

type SetVariable struct {
	ToKey string
	Value interface{}
}

func (b *SetVariable) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *SetVariable) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *SetVariable) kind() string {
	return idSetVariable
}

func (b *SetVariable) Run(ctx context.Context, wce *Environment) (string, Status) {
	wce.setData(b.ToKey, b.Value)

	return log(b, setWorkingData(b.ToKey, fmt.Sprintf("%v", b.Value)), Success)
}
