package netroutine

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
)

const (
	idHashSHA1 = "HashSHA1"
)

func init() {
	blocks[idHashSHA1] = &HashSHA1{}
}

type HashSHA1 struct {
	FromKey string
	ToKey   string
}

func (b *HashSHA1) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *HashSHA1) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *HashSHA1) kind() string {
	return idHashSHA1
}

func (b *HashSHA1) Run(ctx context.Context, wce *Environment) (string, Status) {
	hasher := sha1.New()

	sv, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	s, err := toString(sv)
	if err != nil {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	hasher.Write([]byte(s))

	sha1s := hex.EncodeToString(hasher.Sum(nil))

	wce.setData(b.ToKey, sha1s)

	return log(b, setWorkingData(b.ToKey, sha1s), Success)
}
