package task

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	var taskService TaskService = new(FileWay)
	var task1 Task = taskService.CreateTask("AY-M-ProjectA-20200512-WriteKanbanCode.md")

	if task1.project != "ProjectA" {
		t.Log(task1)
		t.Errorf("Failed")
	}

	if task1.fullName != "AY-M-ProjectA-20200512-WriteKanbanCode.md" {
		t.Log(task1)
		t.Errorf("Failed")
	}
}

func TestIsATask(t *testing.T) {
	var taskService TaskService = new(FileWay)
	var task1 bool = taskService.IsATask("AY-M-ProjectA-20200512-WriteKanbanCode.md")

	if task1 == false {
		t.Log(task1)
		t.Errorf("Failed")
	}

	var task2 = taskService.IsATask("whatever")
	if task2 == true {
		t.Log(task2)
		t.Errorf("Failed")
	}

}

func TestChangeTaskOwner(t *testing.T) {
	var taskService TaskService = new(FileWay)
	var originTask Task = taskService.CreateTask("AY-M-ProjectA-20200512-WriteKanbanCode.md")

	// change owner
	changingTask := ToChangeTask{
		origin:        originTask,
		changeItem:    OWNER,
		changeContent: "WGL",
	}
	var changedTask Task = taskService.ChangeTask(changingTask)

	if changedTask.owner != "WGL" {
		t.Log(changedTask)
		t.Errorf("Failed")
	}

}

func TestChangeTaskpriority(t *testing.T) {
	var taskService TaskService = new(FileWay)

	var originTask Task = taskService.CreateTask("AY-M-ProjectA-20200512-WriteKanbanCode.md")

	// change priority
	var changingTask = ToChangeTask{
		origin:        originTask,
		changeItem:    PRIORITY,
		changeContent: "H",
	}
	var changedTask = taskService.ChangeTask(changingTask)

	if changedTask.priority != "H" {
		t.Log(changedTask)
		t.Errorf("Failed")
	}

}

func TestChangeTaskDeadline(t *testing.T) {
	var taskService TaskService = new(FileWay)
	var originTask Task = taskService.CreateTask("AY-M-ProjectA-20200512-WriteKanbanCode.md")

	// change deadline
	var changingTask = ToChangeTask{
		origin:        originTask,
		changeItem:    DEADLINE,
		changeContent: "20200514",
	}
	var changedTask = taskService.ChangeTask(changingTask)

	if changedTask.deadline != "20200514" {
		t.Log(changedTask)
		t.Errorf("Failed")
	}
}

func TestGetTaskDesc(t *testing.T) {
	var taskService TaskService = new(FileWay)
	var task1 Task = taskService.CreateTask("AY-M-ProjectA-20200512-WriteKanbanCode.md")

	var result = taskService.GetTaskDesc(task1, UNKNOWN)
	if result != "AY-M-ProjectA-20200512-WriteKanbanCode.md" {
		t.Log(result)
		t.Errorf("Failed")
	}
}

func TestFillBlank(t *testing.T) {
	var taskService TaskService = new(FileWay)
	var taskDesc = "1234"

	var result string

	result = taskService.FillBlank(taskDesc, 6)
	if result != "1234  " {
		t.Log(result)
		t.Errorf("Failed")
	}

	result = taskService.FillBlank(taskDesc, 9)
	if result != "1234     " {
		t.Log(result)
		t.Errorf("Failed")
	}
}
