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

	item = getTaskItem("owner")
	assertOwner(item, t)

	item = getTaskItem("OwnEr")
	assertOwner(item, t)

	item = getTaskItem("OWNER")
	assertOwner(item, t)

	item = getTaskItem("PRIORITY1")
	if item != UNKNOWN {
		t.Log(item)
		t.Errorf("Failed")
	}

	item = getTaskItem("PRIoRITY")
	assertPriority(item, t)

	item = getTaskItem("DEAdLINE")
	assertDeadline(item, t)

	item = getTaskItem("PROJEcT")
	assertProject(item, t)

}

func TestTaskItemShortCut(t *testing.T) {
	var item TaskItem

	item = getTaskItem("o")
	assertOwner(item, t)

	item = getTaskItem("own")
	assertOwner(item, t)

	item = getTaskItem("prI")
	assertPriority(item, t)

	item = getTaskItem("pi")
	assertPriority(item, t)

	item = getTaskItem("dl")
	assertDeadline(item, t)

	item = getTaskItem("d")
	assertDeadline(item, t)

	item = getTaskItem("pj")
	assertProject(item, t)

	item = getTaskItem("prj")
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
