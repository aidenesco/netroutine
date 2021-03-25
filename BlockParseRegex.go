package netroutine

import (
	"encoding/json"
	"fmt"
	"regexp"
)

const idBlockParseRegex = "BlockParseRegex"

type BlockParseRegex struct {
	Regex    string
	ToKey    string
	Required bool
}

func (b *BlockParseRegex) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockParseRegex) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockParseRegex) kind() string {
	return idBlockParseRegex
}

func (b *BlockParseRegex) Run(wce *Environment) (string, error) {
	body, err := wce.lastResponseBody()
	if err != nil {
		return log(b, fmt.Sprintf("error getting response body - %v", err), Error)
	}

	reg := regexp.MustCompile(b.Regex)
	found := reg.FindString(body)
	if found == "" {
		if b.Required {
			return log(b, fmt.Sprintf("couldn't find variable %v", b.Regex), Fail)
		}
	}

	wce.setData(b.ToKey, found)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, found), Success)
}
