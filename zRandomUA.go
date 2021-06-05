package netroutine

import (
	"context"
	"encoding/json"

	"github.com/aidenesco/randomua"
)

const idRandomUA = "RandomUA"

func init() {
	blocks[idRandomUA] = &RandomUA{}
}

type RandomUA struct {
	ToKey string
}

func (b *RandomUA) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *RandomUA) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *RandomUA) kind() string {
	return idRandomUA
}

func (b *RandomUA) Run(ctx context.Context, wce *Environment) (string, Status) {

	uas := randomua.Random()

	wce.setData(b.ToKey, uas)

	return log(b, setWorkingData(b.ToKey, uas), Success)
}
