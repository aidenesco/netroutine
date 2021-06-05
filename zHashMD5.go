package netroutine

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

const (
	idHashMD5 = "HashMD5"
)

func init() {
	blocks[idHashMD5] = &HashMD5{}
}

type HashMD5 struct {
	FromKey string
	ToKey   string
}

func (b *HashMD5) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *HashMD5) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *HashMD5) kind() string {
	return idHashMD5
}

func (b *HashMD5) Run(ctx context.Context, wce *Environment) (string, Status) {
	hasher := md5.New()

	s, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	sv, err := toString(s)
	if err != nil {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	hasher.Write([]byte(sv))

	md5s := hex.EncodeToString(hasher.Sum(nil))

	wce.setData(b.ToKey, md5s)

	return log(b, setWorkingData(b.ToKey, md5s), Success)
}
