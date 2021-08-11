package netroutine

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
)

const idParseRegex = "ParseRegex"

func init() {
	blocks[idParseRegex] = &ParseRegex{}
}

type ParseRegex struct {
	Regex    string
	ToKey    string
	Required bool
}

func (b *ParseRegex) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *ParseRegex) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *ParseRegex) kind() string {
	return idParseRegex
}

func (b *ParseRegex) Run(ctx context.Context, wce *Environment) (string, Status) {
	reg := regexp.MustCompile(b.Regex)

	found := reg.FindString(wce.lastResponseBody)
	if found == "" {
		if b.Required {
			return log(b, fmt.Sprintf("couldn't match \"%v\"", b.Regex), Fail)
		}
	}

	wce.setData(b.ToKey, found)

	return log(b, setWorkingData(b.ToKey, found), Success)
}
