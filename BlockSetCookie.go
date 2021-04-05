package netroutine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const idBlockSetCookie = "BlockSetCookie"

type BlockSetCookie struct {
	URL       string
	Name      string
	Value     string
	Variables []string
	Complex   bool
}

func (b *BlockSetCookie) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockSetCookie) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockSetCookie) kind() string {
	return idBlockSetCookie
}

func (b *BlockSetCookie) Run(wce *Environment) (string, Status) {
	purl, err := url.Parse(b.URL)
	if err != nil {
		return log(b, "unable to parse url", Error)
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
		return log(b, fmt.Sprintf("set cookie %v to %v", b.Name, b.Value), Success)
	}

	var sub []interface{}
	for _, va := range b.Variables {
		sv, ok := wce.getData(va)
		if !ok {
			return log(b, fmt.Sprintf("Failed to find \"%v\" variable", va), Error)
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

	return log(b, fmt.Sprintf("set cookie %v to %v", b.Name, value), Success)
}
