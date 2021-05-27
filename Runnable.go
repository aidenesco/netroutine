package netroutine

import (
	"context"
)

var blocks = make(map[string]Runnable)

type Runnable interface {
	Run(ctx context.Context, wce *Environment) (message string, status Status)
	kind() string
	toBytes() ([]byte, error)
	fromBytes([]byte) error
}
