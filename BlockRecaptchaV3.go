package netroutine

import (
	"encoding/json"
	"fmt"
	"github.com/aidenesco/anticaptcha"
)

const idBlockRecaptchaV3 = "RecaptchaV3"

type BlockRecaptchaV3 struct {
	SiteURL  string
	Sitekey  string
	MinScore float64
	ToKey    string
}

func (b *BlockRecaptchaV3) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockRecaptchaV3) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockRecaptchaV3) kind() string {
	return idBlockRecaptchaV3
}

func (b *BlockRecaptchaV3) Run(wce *Environment) (string, error) {
	key, found := wce.getSecret("anticaptcha")
	if !found {
		return log(b, "unable to find anticaptcha api key", Error)
	}

	client := anticaptcha.NewClient(key)

	solution, err := client.RecaptchaV3(b.SiteURL, b.Sitekey, b.MinScore)
	if err != nil {
		return log(b, fmt.Sprintf("error getting recaptcha solution - %v", err), Error)
	}

	wce.setData(b.ToKey, solution.GRecaptchaResponse)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, solution.GRecaptchaResponse), Success)
}
