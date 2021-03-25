package netroutine

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

const idBlockMD5Hash = "BlockMD5Hash"

type BlockMD5Hash struct {
	FromKey string
	ToKey   string
}

func (b *BlockMD5Hash) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockMD5Hash) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockMD5Hash) kind() string {
	return idBlockMD5Hash
}

func (b *BlockMD5Hash) Run(wce *Environment) (string, error) {
	hasher := md5.New()

	s, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "couldn't find the variable to hash", Error)
	}

	sv, err := toString(s)
	if err != nil {
		return log(b, "unable to convert variable to string", Error)
	}

	hasher.Write([]byte(sv))

	md5s := hex.EncodeToString(hasher.Sum(nil))

	wce.setData(b.ToKey, md5s)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, md5s), Success)
}
