package netroutine

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

const idBlockBasicAuth = "BasicAuth"

type BlockBasicAuth struct {
	ToKey       string
	UsernameVar string
	PasswordVar string
}

func (b *BlockBasicAuth) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockBasicAuth) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockBasicAuth) kind() string {
	return idBlockBasicAuth
}

func (b *BlockBasicAuth) Run(wce *Environment) (string, error) {
	u, ok := wce.getData(b.UsernameVar)
	if !ok {
		return log(b, "unable to find username variable", Error)
	}

	us, err := toString(u)
	if err != nil {
		return log(b, "unable to convert variable to string", Error)
	}

	p, ok := wce.getData(b.PasswordVar)
	if !ok {
		return log(b, "unable to find password variable", Error)
	}

	ps, err := toString(p)
	if err != nil {
		return log(b, "unable to convert variable to string", Error)
	}

	built := fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(us+":"+ps)))

	wce.setData(b.ToKey, built)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, built), Success)
}
