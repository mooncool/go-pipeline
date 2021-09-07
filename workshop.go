// @author moon.ghost.jian@gmail.com
package pipeline

type Workshop interface{}

type workshopInstance struct {
	Tasks []Task // tasks in the workshop
}

// NewDefaultWorkshop workshop instantiation
func NewDefaultWorkshop() Workshop {
	instance := new(workshopInstance)

	return instance
}
