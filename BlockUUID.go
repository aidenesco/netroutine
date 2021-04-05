package netroutine

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

const idBlockUUID = "BlockUUID"

type BlockUUID struct {
	ToKey string
}

func (b *BlockUUID) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockUUID) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockUUID) kind() string {
	return idBlockUUID
}

func (b *BlockUUID) Run(wce *Environment) (string, Status) {

	s := uuid.New().String()

	wce.setData(b.ToKey, s)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, s), Success)
}
