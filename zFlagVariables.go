package netroutine

import (
	"context"
	"encoding/json"
)

const idFlagVariables = "FlagVariables"

func init() {
	blocks[idFlagVariables] = &FlagVariables{}
}

type FlagVariables struct {
	Vars  []string
	ToKey string
}

func (b *FlagVariables) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *FlagVariables) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *FlagVariables) kind() string {
	return idFlagVariables
}

func (b *FlagVariables) Run(ctx context.Context, wce *Environment) (string, Status) {
	for _, s := range b.Vars {
		_, ok := wce.getData(s)
		if !ok {
			wce.setExportData(b.ToKey, false)
			return log(b, "set false flag", Success)
		}
	}

	wce.setExportData(b.ToKey, true)

	return log(b, "set true flag", Success)
}
