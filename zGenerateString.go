package netroutine

import (
	"context"
	"encoding/json"
	"math/rand"
	"strings"
	"time"
)

var upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var lower = "abcdefghijklmnopqrstuvwxyz"
var digit = "0123456789"
var symbol = "+/"

const idGenerateString = "GenerateString"

func init() {
	blocks[idGenerateString] = &GenerateString{}
}

type GenerateString struct {
	Base  string
	ToKey string
}

func (b *GenerateString) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *GenerateString) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *GenerateString) kind() string {
	return idGenerateString
}

func (b *GenerateString) Run(ctx context.Context, wce *Environment) (string, Status) {
	built := b.Base

	rand.Seed(time.Now().UnixNano())

	if strings.Contains(built, "~u") {
		for {
			if strings.Contains(built, "~u") {
				built = strings.Replace(built, "~u", randomUpper(), 1)
			} else {
				break
			}
		}
	}

	if strings.Contains(built, "~l") {
		for {
			if strings.Contains(built, "~l") {
				built = strings.Replace(built, "~l", randomLower(), 1)
			} else {
				break
			}
		}
	}

	if strings.Contains(built, "~d") {
		for {
			if strings.Contains(built, "~d") {
				built = strings.Replace(built, "~d", randomDigit(), 1)
			} else {
				break
			}
		}
	}

	if strings.Contains(built, "~s") {
		for {
			if strings.Contains(built, "~s") {
				built = strings.Replace(built, "~s", randomSymbol(), 1)
			} else {
				break
			}
		}
	}

	if strings.Contains(built, "~a") {
		for {
			if strings.Contains(built, "~a") {
				built = strings.Replace(built, "~a", randomAny(), 1)
			} else {
				break
			}
		}
	}

	wce.setData(b.ToKey, built)

	return log(b, setWorkingData(b.ToKey, built), Success)
}

func randomUpper() string {
	return string(upper[rand.Intn(len(upper))])
}

func randomDigit() string {
	return string(digit[rand.Intn(len(digit))])
}

func randomLower() string {
	return string(lower[rand.Intn(len(lower))])
}

func randomSymbol() string {
	return string(symbol[rand.Intn(len(symbol))])
}

func randomAny() string {
	sel := rand.Intn(4)

	switch sel {
	case 0:
		return randomUpper()
	case 1:
		return randomLower()
	case 2:
		return randomDigit()
	case 3:
		return randomSymbol()
	default:
		return randomAny()
	}
}
