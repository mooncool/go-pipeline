package pipeline

import "context"

type pipeSlice struct {
	headerIndex int // queue header index
	rearIndex   int // queue rear index
	tasks       []Task
	workers     []Worker
}

func (p *pipeSlice) queue(ctx context.Context, task Task) error {
	// if len(p.tasks) <= 0 {
	// 	return ErrPipeOverflow
	// }

	// currentIndex := p.headerIndex + 1
	// if currentIndex >= len(p.tasks) {
	// 	currentIndex = 0
	// }
	// p.tasks[currentIndex] = task
	p.tasks = append(p.tasks, task)

	return nil
}

func (p *pipeSlice) dequeue(ctx context.Context) (Task, error) {
	if len(p.tasks) <= 0 {
		return nil, ErrPipeOverflow
	}

	result := p.tasks[0:1]
	_ = copy(p.tasks, p.tasks[1:])

	return result[0], nil
}

func NewPipeSlice(maxSize int, workerCount int) Pipe {
	instance := new(pipeSlice)

	instance.tasks = make([]Task, maxSize)
	instance.headerIndex = 0
	instance.rearIndex = 0
	instance.workers = make([]Worker, workerCount)

	return instance
}
