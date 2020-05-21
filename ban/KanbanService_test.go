package ban

import (
	kt "kanban/task"
	"testing"
)

func TestReadCorrectTask(t *testing.T) {

	var folderPath = ".././unittest/myTasks/02-doing"
	var todoList []string = readFileList(folderPath)
	var todoTasks []kt.Task = readCorrectTasks(todoList)

	var taskNum = len(todoTasks)

	if taskNum != 2 {
		t.Log(todoTasks)
		t.Errorf("Failed")
	}
}

func TestBuildKanban(t *testing.T) {

	var folderPath = ".././unittest/myTasks"

	var kanban Kanban = buildKanban(folderPath)

	if len(kanban.bans) != 4 {
		t.Log(kanban)
		t.Errorf("Failed")
	}
}
