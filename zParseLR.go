package netroutine

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

const idParseLR = "ParseLR"

func init() {
	blocks[idParseLR] = &ParseLR{}
}

type ParseLR struct {
	Left      string
	Right     string
	ToKey     string
	Recursive bool
	Required  bool
}

func (b *ParseLR) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *ParseLR) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *ParseLR) kind() string {
	return idParseLR
}

func (b *ParseLR) Run(ctx context.Context, wce *Environment) (string, Status) {
	if wce.lastResponseBody == "" {
		return log(b, "getting response body", Error)
	}

	if b.Recursive {
		var parsed []string
		var parseFrom string
		parseFrom = wce.lastResponseBody

		for {
			findFirst := strings.Index(parseFrom, b.Left)
			if findFirst == -1 {
				break
			}

			firstIndex := findFirst + len(b.Left)

			findSecond := strings.Index(parseFrom[firstIndex:], b.Right)
			if findSecond == -1 {
				break
			}

			parsedS := parseFrom[firstIndex : firstIndex+findSecond]

			parsed = append(parsed, parsedS)

			parseFrom = parseFrom[firstIndex+findSecond+len(b.Right):]
		}

		if b.Required && len(parsed) == 0 {
			return log(b, "parse is required, but no items were found", Fail)
		}

		wce.setData(b.ToKey, parsed)

		return log(b, setWorkingData(b.ToKey, fmt.Sprintf("%v", parsed)), Success)

	}

	findFirst := strings.Index(wce.lastResponseBody, b.Left)
	if findFirst == -1 {
		if b.Required {
			return log(b, fmt.Sprintf("unable to find required left string in [%s]", base64.StdEncoding.EncodeToString([]byte(wce.lastResponseBody))), Fail)
		}
		return log(b, "unable to find non required left string", Success)
	}

	firstIndex := findFirst + len(b.Left)

	findSecond := strings.Index(wce.lastResponseBody[firstIndex:], b.Right)
	if findSecond == -1 {
		if b.Required {
			return log(b, "unable to find required right string", Fail)
		}
		return log(b, "unable to find non required right string", Success)
	}

	parsed := wce.lastResponseBody[firstIndex : firstIndex+findSecond]

	wce.setData(b.ToKey, parsed)

	return log(b, setWorkingData(b.ToKey, fmt.Sprintf("%v", parsed)), Success)
}
