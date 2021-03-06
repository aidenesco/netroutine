package netroutine

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

const idParseJSON = "ParseJSON"

func init() {
	blocks[idParseJSON] = &ParseJSON{}
}

type ParseJSON struct {
	ToKey     string
	Path      []string
	Recursive bool
	Required  bool
}

func (b *ParseJSON) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *ParseJSON) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *ParseJSON) kind() string {
	return idParseJSON
}

func (b *ParseJSON) Run(ctx context.Context, wce *Environment) (string, Status) {
	parsed, err := gabs.ParseJSON([]byte(wce.lastResponseBody))
	if err != nil {
		return log(b, reportError("loading to gabs", err), Error)
	}

	value := parsed.Search(b.Path...)

	if b.Recursive {
		var data []interface{}

		for _, obj := range value.Children() {
			data = append(data, obj.Data())
		}

		if b.Required && len(data) == 0 {
			return log(b, "parse is required, but no items were found", Fail)
		}

		wce.setData(b.ToKey, data)

		return log(b, setWorkingData(b.ToKey, fmt.Sprintf("%v", data)), Success)
	}

	if value == nil && b.Required {
		return log(b, fmt.Sprintf("couldn't find required variable \"%v\"", b.Path), Fail)
	}

	data := value.Data()

	if data == nil {
		data = ""
	}

	wce.setData(b.ToKey, data)

	return log(b, setWorkingData(b.ToKey, fmt.Sprintf("%v", data)), Success)
}
