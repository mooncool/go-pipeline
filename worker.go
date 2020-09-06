// @author moon.ghost.jian@gmail.com
package pipeline

import "context"

type Worker interface {
	process(ctx context.Context) error
	putTo(pipe Pipe) error
	fetchFrom(pipe Pipe) ([]byte, error)
}
