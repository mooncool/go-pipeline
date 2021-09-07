// @author moon.ghost.jian@gmail.com
package pipeline

import "context"

type Worker interface {
	Process(ctx context.Context, task Task) error // Core process function

	// putTo(pipe Pipe) error
	// fetchFrom(pipe Pipe) ([]byte, error)
}

type workerInstance struct {
	isRunning bool
	thread    chan Task
}

func (w *workerInstance) Process(ctx context.Context, task Task) error {
	return nil
}

// NewWorker worker instantiation
func NewWorker() Worker {
	instance := new(workerInstance)

	return instance
}
