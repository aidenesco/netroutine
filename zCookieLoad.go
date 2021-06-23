package netroutine

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func init() {
	blocks[idLoadCookies] = &LoadCookies{}
}

const idLoadCookies = "LoadCookies"

type LoadCookies struct {
	URL     string
	FromKey string
}

func (b *LoadCookies) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *LoadCookies) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *LoadCookies) kind() string {
	return idLoadCookies
}

func (b *LoadCookies) Run(ctx context.Context, wce *Environment) (string, Status) {
	purl, err := url.Parse(b.URL)
	if err != nil {
		return log(b, reportError("parsing url", err), Error)
	}

	data, found := wce.getData(b.FromKey)
	if !found {
		return log(b, missingWorkingData(b.FromKey), Error)
	}

	bData, ok := data.([]byte)
	if !ok {
		return log(b, reportWrongType(b.FromKey), Error)
	}

	var cookies []*http.Cookie

	err = json.Unmarshal(bData, &cookies)
	if err != nil {
		return log(b, reportError("unmarshaling", err), Error)
	}

	wce.Client.Jar.SetCookies(purl, cookies)

	return log(b, fmt.Sprintf("loaded cookies: %s", bData), Success)
}
