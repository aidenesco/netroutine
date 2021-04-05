package netroutine

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

const idBlockRandomChoiceList = "BlockRandomChoiceList"

type BlockRandomChoiceList struct {
	Choices []string
	ToKey   string
}

func (b *BlockRandomChoiceList) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockRandomChoiceList) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockRandomChoiceList) kind() string {
	return idBlockRandomChoiceList
}

func (b *BlockRandomChoiceList) Run(wce *Environment) (string, Status) {
	rand.Seed(time.Now().UnixNano())

	choice := b.Choices[rand.Intn(len(b.Choices))]

	wce.setData(b.ToKey, choice)

	return log(b, fmt.Sprintf("set random choice from list - %v", choice), Success)
}
