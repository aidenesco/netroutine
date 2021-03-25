package netroutine

import (
	"encoding/json"
	"fmt"
	"github.com/aidenesco/anticaptcha"
)

const idBlockRecaptcha = "BlockRecaptcha"

type BlockRecaptcha struct {
	SiteURL string
	Sitekey string
	ToKey   string
}

func (b *BlockRecaptcha) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockRecaptcha) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockRecaptcha) kind() string {
	return idBlockRecaptcha
}

func (b *BlockRecaptcha) Run(wce *Environment) (string, error) {

	key, found := wce.getSecret("anticaptcha")
	if !found {
		return log(b, "unable to find anticaptcha api key", Error)
	}

	client := anticaptcha.NewClient(key)

	solution, err := client.RecaptchaProxyless(b.SiteURL, b.Sitekey)
	if err != nil {
		return log(b, fmt.Sprintf("error getting recaptcha solution - %v", err), Error)
	}

	wce.setData(b.ToKey, solution.GRecaptchaResponse)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, solution.GRecaptchaResponse), Success)
}
