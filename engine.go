// Core engine
package pipeline

type Engine interface {
	Start() error // start the engine
	Pause() error // pause the engine
	Stop() error  // stop the engine

	Bind(pipe Pipe, worker Worker, concurrency int) // register a new worker
}

type pipelineEngine struct {
	workers []Worker
}

func (engine *pipelineEngine) Start() error {
	return nil
}

func (engine *pipelineEngine) Pause() error {
	return nil
}

func (engine *pipelineEngine) Stop() error {
	return nil
}

func (engine *pipelineEngine) Bind(pipe Pipe, worker Worker, concurrency int) {
	engine.workers = append(engine.workers, worker)
}

// NewDefaultEngine create a default pipeline engine instance
func NewDefaultEngine() Engine {
	engine := new(pipelineEngine)

	return engine
}
