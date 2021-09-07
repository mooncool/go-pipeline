// @author moon.ghost.jian@gmail.com
package pipeline

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewTask(t *testing.T) {
	Convey("test string task", t, func() {
		taskContent := "string task"
		t := NewTask([]byte(taskContent))
		So(t.GetID(), ShouldBeGreaterThan, 0)
		So(string(t.GetData()), ShouldEqual, taskContent)
	})
	Convey("test int task", t, func() {
		// taskContent := 123
		t := NewTask([]byte{1, 2})
		So(t.GetID(), ShouldBeGreaterThan, 0)
		data := t.GetData()
		So(bytes.Equal(data, []byte{1, 2}), ShouldEqual, true)
	})
}
