package netroutine

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

const idBlockSHA1Hash = "BlockSHA1Hash"

type BlockSHA1Hash struct {
	FromKey string
	ToKey   string
}

func (b *BlockSHA1Hash) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockSHA1Hash) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockSHA1Hash) kind() string {
	return idBlockSHA1Hash
}

func (b *BlockSHA1Hash) Run(wce *Environment) (string, error) {
	hasher := sha1.New()

	sv, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "couldn't find variable to hash", Error)
	}

	s, err := toString(sv)
	if err != nil {
		return log(b, "unable to convert variable to string", Error)
	}

	hasher.Write([]byte(s))

	sha1s := hex.EncodeToString(hasher.Sum(nil))

	wce.setData(b.ToKey, sha1s)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, sha1s), Success)
}
