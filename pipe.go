// Core engine
package pipeline

type Pipe interface {
	enqueue(task []byte) error
	dequeue() ([]byte, error)
	size() (int, error)
}
