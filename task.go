// @author moon.ghost.jian@gmail.com
package pipeline

import "time"

// Task task interface
type Task interface {
	GetID() int64
	GetName() string
	GetData() []byte
}

type taskInstance struct {
	id   int64  // unique ID
	name string // unique task name
	data []byte // task content data
}

func (*taskInstance) genID() int64 {
	return time.Now().Unix()
}

func (t *taskInstance) GetID() int64 {
	return t.id
}

func (t *taskInstance) GetName() string {
	return t.name
}

func (t *taskInstance) GetData() []byte {
	return t.data
}

func NewTask(data []byte) Task {
	instance := new(taskInstance)

	instance.id = instance.genID()
	instance.data = data

	return instance
}
