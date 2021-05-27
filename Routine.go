package netroutine

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"time"
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

func (r *Routine) Run(ctx context.Context, wce *Environment) {
	for _, v := range r.blocks {
		attempts := 0
		for {
			if attempts >= wce.maxRetry && wce.maxRetry != -1 {
				return
			}

			msg, status := v.Run(ctx, wce)

			wce.addLog(msg)
			wce.Status = status

			switch status {
			case Error:
				return
			case Retry:
				wce.Client.CloseIdleConnections()

				//Retries are often network failures or rate limiting
				time.Sleep(wce.retrySleep)

				attempts++
				continue
			case Fail:
				return
			case Custom:
				return
			case Success:
			}
			break
		}
	}
}

func NewRoutine(b ...Runnable) *Routine {
	return &Routine{blocks: b}
}
