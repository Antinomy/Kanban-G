package task

import (
	"testing"
)

func TestToTask(t *testing.T) {

    var taskName string = "AY-M-ProjectA-20200512-WriteKanbanCode.md"
    var task1 Task = toTask(taskName)
    t.Log(task1)

    if task1.project != "ProjectA" {
        t.Errorf("Failed")
    }
}

func TestIsATask(t *testing.T) {
    var taskName string = "AY-M-ProjectA-20200512-WriteKanbanCode.md"
    var task1 bool = isATask(taskName)

    if task1 == false {
        t.Errorf("Failed")
    }

    taskName = "whatever"
    var task2 = isATask(taskName)
    if task2 == true {
        t.Errorf("Failed")
    }

}
