package netroutine

import (
	"context"
	"encoding/json"
)

const idFlagToStatus = "FlagToStatus"

func init() {
	blocks[idFlagToStatus] = &FlagToStatus{}
}

type FlagToStatus struct {
	FromKey string
	IfTrue  Status
	IfFalse Status
}

func (b *FlagToStatus) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *FlagToStatus) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *FlagToStatus) kind() string {
	return idFlagToStatus
}

func (b *FlagToStatus) Run(ctx context.Context, wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	bo, ok := v.(bool)
	if !ok {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	if bo {
		return log(b, "flag is true", b.IfTrue)
	} else {
		return log(b, "flag is false", b.IfFalse)
	}
}
