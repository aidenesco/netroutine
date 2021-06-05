package netroutine

import (
	"context"
	"encoding/json"
	"net/url"
)

const idParseCookies = "ParseCookies"

func init() {
	blocks[idParseCookies] = &ParseCookies{}
}

type ParseCookies struct {
	URL        string
	CookieName string
	ToKey      string
	Required   bool
}

func (b *ParseCookies) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *ParseCookies) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *ParseCookies) kind() string {
	return idParseCookies
}

func (b *ParseCookies) Run(ctx context.Context, wce *Environment) (string, Status) {
	parsed, err := url.Parse(b.URL)
	if err != nil {
		return log(b, reportError("parsing url", err), Error)
	}

	cookies := wce.Client.Jar.Cookies(parsed)

	for _, v := range cookies {
		if v.Name == b.CookieName {
			wce.setData(b.ToKey, v.Value)
			return log(b, setWorkingData(b.ToKey, v.Value), Success)
		}
	}

	if b.Required {
		return log(b, "unable to find required cookie", Fail)
	}

	return log(b, "unable to find non required cookie", Success)
}
