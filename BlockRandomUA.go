package netroutine

import (
	"encoding/json"
	"fmt"
	"github.com/aidenesco/randomua"
)

const idBlockRandomUA = "BlockRandomUA"

type BlockRandomUA struct {
	ToKey string
}

func (b *BlockRandomUA) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockRandomUA) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockRandomUA) kind() string {
	return idBlockRandomUA
}

func (b *BlockRandomUA) Run(wce *Environment) (string, error) {

	uas := randomua.Random()

	wce.setData(b.ToKey, uas)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, uas), Success)
}
