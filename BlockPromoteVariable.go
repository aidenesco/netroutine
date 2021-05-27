package netroutine

import (
	"context"
	"encoding/json"
	"fmt"
)

func init() {
	blocks[idPromoteVariable] = &PromoteVariable{}
}

const idPromoteVariable = "PromoteVariable"

type PromoteVariable struct {
	FromKey string
}

func (b *PromoteVariable) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *PromoteVariable) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *PromoteVariable) kind() string {
	return idPromoteVariable
}

func (b *PromoteVariable) Run(ctx context.Context, wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	wce.setExportData(b.FromKey, v)

	return log(b, fmt.Sprintf("promoted the working data at \"%s\" to export data", b.FromKey), Success)
}
