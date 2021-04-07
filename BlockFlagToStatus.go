package netroutine

import (
	"encoding/json"
)

const idBlockFlagToStatus = "BlockFlagToStatus"

type BlockFlagToStatus struct {
	FromKey string
	IfTrue  Status
	IfFalse Status
}

func (b *BlockFlagToStatus) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockFlagToStatus) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockFlagToStatus) kind() string {
	return idBlockFlagToStatus
}

func (b *BlockFlagToStatus) Run(wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "variable not found", Error)
	}

	bo, ok := v.(bool)
	if !ok {
		return log(b, "variable not a boolean", Error)
	}

	if bo {
		return log(b, "flag is true", b.IfTrue)
	} else {
		return log(b, "flag is false", b.IfFalse)
	}
}
