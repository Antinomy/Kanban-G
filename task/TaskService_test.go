package task

import (
	"testing"
)

func TestToTask(t *testing.T) {

	var fileName string = "AY-M-ProjectA-20200512-WriteKanbanCode.md"

	var task1 Task
	task1 = toTask(fileName)
	t.Log(task1)

	if task1.project != "ProjectA" {
		t.Errorf("Failed")
	}
}

func TestIsATask(t *testing.T) {
	var fileName string = "AY-M-ProjectA-20200512-WriteKanbanCode.md"

	var task1 bool
	task1 = isATask(fileName)

	if task1 == false {
		t.Errorf("Failed")
	}

	fileName = "whatever"
	var task2 = isATask(fileName)
	if task2 == true {
		t.Errorf("Failed")
	}

}
