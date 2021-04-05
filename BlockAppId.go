package netroutine

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

const idBlockAppId = "AppId"

type BlockAppId struct {
	ToKey string
}

func (b *BlockAppId) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockAppId) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockAppId) kind() string {
	return idBlockAppId
}

func (b *BlockAppId) Run(wce *Environment) (string, Status) {

	s := strings.ToUpper(uuid.New().String())

	wce.setData(b.ToKey, s)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, s), Success)
}
