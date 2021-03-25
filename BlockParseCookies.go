package netroutine

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const idBlockParseCookies = "BlockParseCookies"

type BlockParseCookies struct {
	URL        string
	CookieName string
	ToKey      string
	Required   bool
}

func (b *BlockParseCookies) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockParseCookies) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockParseCookies) kind() string {
	return idBlockParseCookies
}

func (b *BlockParseCookies) Run(wce *Environment) (string, error) {
	parsed, err := url.Parse(b.URL)
	if err != nil {
		return log(b, fmt.Sprintf("error parsing url - %v", err), Error)
	}

	cookies := wce.Client.Jar.Cookies(parsed)

	for _, v := range cookies {
		if v.Name == b.CookieName {
			wce.setData(b.ToKey, v.Value)
			return log(b, fmt.Sprintf("set %v to %v", b.ToKey, v.Value), Success)
		}
	}

	if b.Required {
		return log(b, "unable to find required cookie", Fail)
	}

	return log(b, "unable to find non required cookie", Success)
}
