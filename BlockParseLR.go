package netroutine

import (
	"encoding/json"
	"fmt"
	"strings"
)

const idBlockParseLR = "BlockParseLR"

type BlockParseLR struct {
	Left      string
	Right     string
	ToKey     string
	Recursive bool
	Required  bool
}

func (b *BlockParseLR) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockParseLR) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockParseLR) kind() string {
	return idBlockParseLR
}

func (b *BlockParseLR) Run(wce *Environment) (string, error) {
	strbody, err := wce.lastResponseBody()
	if err != nil {
		return log(b, fmt.Sprintf("error getting response body - %v", err), Error)
	}

	if b.Recursive {
		var parsed []string
		var parseFrom string
		parseFrom = strbody

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

		return log(b, fmt.Sprintf("set %v to %v", b.ToKey, parsed), Success)

	}

	findFirst := strings.Index(strbody, b.Left)
	if findFirst == -1 {
		if b.Required {
			return log(b, "unable to find required left string", Fail)
		}
		return log(b, "unable to find non required left string", Success)
	}

	firstIndex := findFirst + len(b.Left)

	findSecond := strings.Index(strbody[firstIndex:], b.Right)
	if findSecond == -1 {
		if b.Required {
			return log(b, "unable to find required right string", Fail)
		}
		return log(b, "unable to find non required right string", Success)
	}

	parsed := strbody[firstIndex : firstIndex+findSecond]

	wce.setData(b.ToKey, parsed)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, parsed), Success)
}
