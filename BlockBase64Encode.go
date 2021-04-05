package netroutine

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

const idBlockBase64Encode = "Base64Encode"

type BlockBase64Encode struct {
	FromKey string
	ToKey   string
}

func (b *BlockBase64Encode) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockBase64Encode) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockBase64Encode) kind() string {
	return idBlockBase64Encode
}

func (b *BlockBase64Encode) Run(wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "missing input variable", Error)
	}

	s, err := toString(v)
	if err != nil {
		return log(b, "unable to convert variable to string", Error)
	}

	e := base64.StdEncoding.EncodeToString([]byte(s))

	wce.setData(b.ToKey, e)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, e), Success)
}
