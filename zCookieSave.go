package netroutine

import (
	"context"
	"encoding/json"
	"net/url"
)

func init() {
	blocks[idSaveCookies] = &SaveCookies{}
}

const idSaveCookies = "SaveCookies"

type SaveCookies struct {
	URL   string
	ToKey string
}

func (b *SaveCookies) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *SaveCookies) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *SaveCookies) kind() string {
	return idSaveCookies
}

func (b *SaveCookies) Run(ctx context.Context, wce *Environment) (string, Status) {
	purl, err := url.Parse(b.URL)
	if err != nil {
		return log(b, reportError("parsing url", err), Error)
	}

	cookies := wce.Client.Jar.Cookies(purl)

	data, err := json.Marshal(cookies)
	if err != nil {
		return log(b, reportError("marshaling cookies", err), Error)
	}

	wce.setData(b.ToKey, data)

	return log(b, setWorkingData(b.ToKey, data), Success)
}
