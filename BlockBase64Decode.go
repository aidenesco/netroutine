package netroutine

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

const idBlockBase64Decode = "Base64Decode"

type BlockBase64Decode struct {
	FromKey string
	ToKey   string
}

func (b *BlockBase64Decode) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockBase64Decode) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockBase64Decode) kind() string {
	return idBlockBase64Decode
}

func (b *BlockBase64Decode) Run(wce *Environment) (string, Status) {
	v, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "missing input variable", Error)
	}

	s, err := toString(v)
	if err != nil {
		return log(b, "unable to convert variable to string", Error)
	}

	e, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return log(b, fmt.Sprintf("error decoding - %v", err), Error)
	}

	es := string(e)

	wce.setData(b.ToKey, es)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, es), Success)
}
