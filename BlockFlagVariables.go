package netroutine

import (
	"encoding/json"
)

const idBlockFlagVariables = "BlockFlagVariables"

type BlockFlagVariables struct {
	Vars  []string
	ToKey string
}

func (b *BlockFlagVariables) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockFlagVariables) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockFlagVariables) kind() string {
	return idBlockFlagVariables
}

func (b *BlockFlagVariables) Run(wce *Environment) (string, error) {
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
