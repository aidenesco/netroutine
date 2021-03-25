package netroutine

import (
	"encoding/json"
	"fmt"
)

const idBlockVarToCap = "BlockVarToCap"

type BlockVarToCap struct {
	Key string
}

func (b *BlockVarToCap) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockVarToCap) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockVarToCap) kind() string {
	return idBlockVarToCap
}

func (b *BlockVarToCap) Run(wce *Environment) (string, error) {
	v, ok := wce.getData(b.Key)
	if !ok {
		return log(b, "couldn't find variable to promote", Error)
	}

	wce.setExportData(b.Key, v)

	return log(b, fmt.Sprintf("promoted %v to capture", b.Key), Success)
}
