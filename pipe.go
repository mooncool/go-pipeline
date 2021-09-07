// Core engine
package pipeline

import (
	"context"
	"errors"
)

var (
	ErrPipeOverflow = errors.New("pipe capacity is not enough")
)

type Pipe interface {
	queue(ctx context.Context, task Task) error
	dequeue(ctx context.Context) (Task, error)
	// size() (int, error)
}

// func(p *pipeInstance) queue(ctx context.Context, task Task) error

func CreatePipe(impl string, workerCount, maxSize int) Pipe {
	var pipe Pipe

	switch impl {
	case "slice":
		pipe = NewPipeSlice(maxSize, workerCount)
	default:
		pipe = NewPipeSlice(100, 1)
	}

	return pipe
}
