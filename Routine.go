package netroutine

import (
	"encoding/json"
	"errors"
	"reflect"
)

type Routine struct {
	blocks []Runnable
}

func RoutineFromBytes(raw []byte) (*Routine, error) {
	var toBuild []struct {
		Kind string
		Data []byte
	}
	var builtBlocks []Runnable

	if err := json.Unmarshal(raw, &toBuild); err != nil {
		return nil, err
	}

	for _, b := range toBuild {
		foundType, ok := blocks[b.Kind]
		if !ok {
			return nil, errors.New("unknown block type:" + b.Kind)
		}

		block, ok := reflect.New(reflect.TypeOf(foundType)).Interface().(Runnable)
		if !ok {
			return nil, errors.New("block does not conform to the Runnable interface")
		}

		if err := block.fromBytes(b.Data); err != nil {
			return nil, err
		}

		builtBlocks = append(builtBlocks, block)
	}

	return &Routine{blocks: builtBlocks}, nil
}

func (r *Routine) ToBytes() ([]byte, error) {
	var blocks []struct {
		Kind string
		Data []byte
	}

	for _, b := range r.blocks {
		data, err := b.toBytes()
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, struct {
			Kind string
			Data []byte
		}{Kind: b.kind(), Data: data})
	}
	return json.Marshal(blocks)
}

func NewRoutine(b ...Runnable) *Routine {
	return &Routine{blocks: b}
}
