package netroutine

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func init() {
	blocks[idSetCookie] = &SetCookie{}
}

const idSetCookie = "SetCookie"

type SetCookie struct {
	URL       string
	Name      string
	Value     string
	Variables []string
	Complex   bool
}

func (b *SetCookie) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *SetCookie) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *SetCookie) kind() string {
	return idSetCookie
}

func (b *SetCookie) Run(ctx context.Context, wce *Environment) (string, Status) {
	purl, err := url.Parse(b.URL)
	if err != nil {
		return log(b, reportError("parsing url", err), Error)
	}

	if !b.Complex {
		wce.Client.Jar.SetCookies(purl, []*http.Cookie{
			{
				Name:   b.Name,
				Value:  b.Value,
				Raw:    b.Name + "=" + b.Value,
				MaxAge: 3600,
			},
		})
		return log(b, fmt.Sprintf("set cookie \"%v\" to \"%v\"", b.Name, b.Value), Success)
	}

	var sub []interface{}
	for _, va := range b.Variables {
		sv, ok := wce.getData(va)
		if !ok {
			return log(b, missingWorkingData(va), Error)
		}
		sub = append(sub, sv)
	}

	value := fmt.Sprintf(b.Value, sub...)

	wce.Client.Jar.SetCookies(purl, []*http.Cookie{
		{
			Name:   b.Name,
			Value:  value,
			Raw:    b.Name + "=" + value,
			MaxAge: 3600,
		},
	})

	return log(b, fmt.Sprintf("set cookie \"%v\" to \"%v\"", b.Name, value), Success)
}
