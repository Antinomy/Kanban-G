package task

import (
	"testing"
)

func TestTask(t *testing.T) {
	var task1 Task

	task1.Owner = "AY"
	task1.Priority = "M"
	task1.Project = "ProjectA"
	task1.Deadline = "20200512"
	task1.tittle = "WriteKanbanCode.md"

	if task1.Owner != "AY" {
		t.Log(task1)
		t.Errorf("Failed")
	}

}

func TestTaskItem(t *testing.T) {
	var item TaskItem

	item = GetTaskItem("owner")
	assertOwner(item, t)

	item = GetTaskItem("OwnEr")
	assertOwner(item, t)

	item = GetTaskItem("OWNER")
	assertOwner(item, t)

	item = GetTaskItem("PRIORITY1")
	if item != UNKNOWN {
		t.Log(item)
		t.Errorf("Failed")
	}

	item = GetTaskItem("PRIoRITY")
	assertPriority(item, t)

	item = GetTaskItem("DEAdLINE")
	assertDeadline(item, t)

	item = GetTaskItem("PROJEcT")
	assertProject(item, t)

}

func TestTaskItemShortCut(t *testing.T) {
	var item TaskItem

	item = GetTaskItem("o")
	assertOwner(item, t)

	item = GetTaskItem("own")
	assertOwner(item, t)

	item = GetTaskItem("prI")
	assertPriority(item, t)

	item = GetTaskItem("pi")
	assertPriority(item, t)

	item = GetTaskItem("dl")
	assertDeadline(item, t)

	item = GetTaskItem("d")
	assertDeadline(item, t)

	item = GetTaskItem("pj")
	assertProject(item, t)

	item = GetTaskItem("prj")
	assertProject(item, t)

}

func assertOwner(item TaskItem, t *testing.T) {
	if item != OWNER {
		t.Log(item)
		t.Errorf("Failed")
	}
}

func assertPriority(item TaskItem, t *testing.T) {
	if item != PRIORITY {
		t.Log(item)
		t.Errorf("Failed")
	}
}

func assertDeadline(item TaskItem, t *testing.T) {
	if item != DEADLINE {
		t.Log(item)
		t.Errorf("Failed")
	}
}

func assertProject(item TaskItem, t *testing.T) {
	if item != PROJECT {
		t.Log(item)
		t.Errorf("Failed")
	}
}
