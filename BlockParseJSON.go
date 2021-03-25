package netroutine

import (
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs/v2"
)

const idBlockParseJSON = "BlockParseJSON"

type BlockParseJSON struct {
	ToKey     string
	Path      []string
	Recursive bool
	Required  bool
}

func (b *BlockParseJSON) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockParseJSON) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockParseJSON) kind() string {
	return idBlockParseJSON
}

func (b *BlockParseJSON) Run(wce *Environment) (string, error) {

	body, err := wce.lastResponseBody()
	if err != nil {
		return log(b, fmt.Sprintf("error getting response body - %v", err), Error)
	}

	parsed, err := gabs.ParseJSON([]byte(body))
	if err != nil {
		return log(b, fmt.Sprintf("error loading body into gabs - %v", err), Error)
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

		return log(b, fmt.Sprintf("set %v to %v", b.ToKey, data), Success)
	}

	data := value.Data()
	if value == nil && b.Required {
		return log(b, fmt.Sprintf("couldn't find required variable %v", b.Path), Fail)
	}

	if data == nil {
		data = ""
	}

	wce.setData(b.ToKey, data)

	return log(b, fmt.Sprintf("set %v to %v", b.ToKey, data), Success)
}
