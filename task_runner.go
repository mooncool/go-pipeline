// @author moon.ghost.jian@gmail.com
package pipeline

import "log"

type TaskRunner interface {
	Start()
	Stop()
}

type TaskRunnerConfig struct {
	WorkerNum int
}

type defaultTaskRunner struct {
	pipelines []chan interface{}
}

func (t *defaultTaskRunner) Start() {
	for _, p := range t.pipelines {
		go func(pipeline chan interface{}) {
			for {
				task := <-pipeline
				log.Printf("task: %v", task)
			}
		}(p)
	}
}

func (t *defaultTaskRunner) Stop() {}

// NewTaskRunnerWithConfig create task runner ins
func NewTaskRunnerWithConfig(config *TaskRunnerConfig) TaskRunner {
	tr := new(defaultTaskRunner)

	tr.pipelines = make([]chan interface{}, config.WorkerNum)
	for i := 0; i < config.WorkerNum; i++ {
		tr.pipelines[i] = make(chan interface{})
	}

	return tr
}
