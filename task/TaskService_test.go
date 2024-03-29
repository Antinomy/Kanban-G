package task

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	var taskService TaskService = new(FileWay)
	var task1 Task = taskService.CreateTask("AY-M-ProjectA-20200512-WriteKanbanCode.md")

	if task1.Project != "ProjectA" {
		t.Log(task1)
		t.Errorf("Failed")
	}

	if task1.FullName != "AY-M-ProjectA-20200512-WriteKanbanCode.md" {
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
		Origin:        originTask,
		ChangeItem:    OWNER,
		ChangeContent: "WGL",
	}
	var changedTask Task = taskService.ChangeTask(changingTask)

	if changedTask.Owner != "WGL" {
		t.Log(changedTask)
		t.Errorf("Failed")
	}

	if changedTask.FullName != "WGL-M-ProjectA-20200512-WriteKanbanCode.md" {
		t.Log(changedTask)
		t.Errorf("Failed")
	}

}

func TestChangeTaskpriority(t *testing.T) {
	var taskService TaskService = new(FileWay)

	var originTask Task = taskService.CreateTask("AY-M-ProjectA-20200512-WriteKanbanCode.md")

	// change priority
	var changingTask = ToChangeTask{
		Origin:        originTask,
		ChangeItem:    PRIORITY,
		ChangeContent: "H",
	}
	var changedTask = taskService.ChangeTask(changingTask)

	if changedTask.Priority != "H" {
		t.Log(changedTask)
		t.Errorf("Failed")
	}

	if changedTask.FullName != "AY-H-ProjectA-20200512-WriteKanbanCode.md" {
		t.Log(changedTask)
		t.Errorf("Failed")
	}

}

func TestChangeTaskDeadline(t *testing.T) {
	var taskService TaskService = new(FileWay)
	var originTask Task = taskService.CreateTask("AY-M-ProjectA-20200512-WriteKanbanCode.md")

	// change deadline
	var changingTask = ToChangeTask{
		Origin:        originTask,
		ChangeItem:    DEADLINE,
		ChangeContent: "20200514",
	}
	var changedTask = taskService.ChangeTask(changingTask)

	if changedTask.Deadline != "20200514" {
		t.Log(changedTask)
		t.Errorf("Failed")
	}

	if changedTask.FullName != "AY-M-ProjectA-20200514-WriteKanbanCode.md" {
		t.Log(changedTask)
		t.Errorf("Failed")
	}
}

func TestGetTaskDesc(t *testing.T) {
	var taskService TaskService = new(FileWay)
	var task1 Task = taskService.CreateTask("AY-M-ProjectA-20200512-WriteKanbanCode.md")

	var result = taskService.GetTaskDesc(task1, UNKNOWN, false)
	if result != "AY-M-ProjectA-20200512-WriteKanbanCode.md" {
		t.Log(result)
		t.Errorf("Failed")
	}

	task1.Key = "t1"

	result = taskService.GetTaskDesc(task1, UNKNOWN, true)
	if result != "[t1] AY-M-ProjectA-20200512-WriteKanbanCode.md" {
		t.Log(result)
		t.Errorf("Failed")
	}

	result = taskService.GetTaskDesc(task1, OWNER, true)
	if result != "[t1] M-ProjectA-20200512-WriteKanbanCode.md" {
		t.Log(result)
		t.Errorf("Failed")
	}

	result = taskService.GetTaskDesc(task1, PRIORITY, true)
	if result != "[t1] AY-ProjectA-20200512-WriteKanbanCode.md" {
		t.Log(result)
		t.Errorf("Failed")
	}

	result = taskService.GetTaskDesc(task1, PROJECT, true)
	if result != "[t1] AY-M-20200512-WriteKanbanCode.md" {
		t.Log(result)
		t.Errorf("Failed")
	}

	result = taskService.GetTaskDesc(task1, DEADLINE, true)
	if result != "[t1] AY-M-ProjectA-12-WriteKanbanCode.md" {
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
