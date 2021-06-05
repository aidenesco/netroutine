package netroutine

import (
	"context"
	"encoding/json"
	"fmt"
)

const idBuilderString = "BuilderString"

func init() {
	blocks[idBuilderString] = &BuilderString{}
}

type BuilderString struct {
	Variables []string
	Base      string
	ToKey     string
}

func (b *BuilderString) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BuilderString) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BuilderString) kind() string {
	return idBuilderString
}

func (b *BuilderString) Run(ctx context.Context, wce *Environment) (string, Status) {
	var sub []interface{}
	for _, v := range b.Variables {
		sv, ok := wce.getData(v)
		if !ok {
			return log(b, missingWorkingData(v), Error)
		}
		sub = append(sub, sv)
	}

	fs := fmt.Sprintf(b.Base, sub...)

	wce.setData(b.ToKey, fs)

	return log(b, setWorkingData(b.ToKey, fs), Success)
}
