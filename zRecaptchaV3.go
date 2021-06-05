package netroutine

import (
	"context"
	"encoding/json"

	"github.com/aidenesco/anticaptcha"
)

func init() {
	blocks[idRecaptchaV3] = &RecaptchaV3{}
}

const idRecaptchaV3 = "RecaptchaV3"

type RecaptchaV3 struct {
	SiteURL      string
	Sitekey      string
	MinScore     float64
	ToKey        string
	IsEnterprise bool
}

func (b *RecaptchaV3) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *RecaptchaV3) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *RecaptchaV3) kind() string {
	return idRecaptchaV3
}

func (b *RecaptchaV3) Run(ctx context.Context, wce *Environment) (string, Status) {
	key, found := wce.getSecret("anticaptcha")
	if !found {
		return log(b, missingSecret("anticaptcha"), Error)
	}

	client := anticaptcha.NewClient(key)

	if b.IsEnterprise {
		solution, err := client.RecaptchaV3Enterprise(ctx, b.SiteURL, b.Sitekey, b.MinScore)
		if err != nil {
			return log(b, reportError("getting recaptcha token", err), Error)
		}
		wce.setData(b.ToKey, solution.GRecaptchaResponse)

		return log(b, setWorkingData(b.ToKey, solution.GRecaptchaResponse), Success)
	} else {
		solution, err := client.RecaptchaV3Proxyless(ctx, b.SiteURL, b.Sitekey, b.MinScore)
		if err != nil {
			return log(b, reportError("getting recaptcha token", err), Error)
		}

		wce.setData(b.ToKey, solution.GRecaptchaResponse)

		return log(b, setWorkingData(b.ToKey, solution.GRecaptchaResponse), Success)
	}
}
