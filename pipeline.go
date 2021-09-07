// @author moon.ghost.jian@gmail.com
package pipeline

import (
	"context"
	"errors"
)

var (
	ErrPipelineOverflow = errors.New("pipeline capacity is not enough")
)

type Pipeline interface {
	queue(ctx context.Context, task Task) error
	dequeue(ctx context.Context) (Task, error)
}

type pipelineInstance struct {
	workers []Worker
}
