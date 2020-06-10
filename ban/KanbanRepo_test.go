package ban

import (
	kt "kanban/task"
	"testing"
)

func TestReadCorrectTask(t *testing.T) {

	var folderPath = ".././unittest/myTasks/01-todo"
	var todoList []string = readFileList(folderPath)
	var todoTasks []kt.Task = readCorrectTasks(todoList, "t")

	var taskNum = len(todoTasks)

	if taskNum != 2 {
		t.Log(todoTasks)
		t.Errorf("Failed")
	}

	if todoTasks[0].Key != "t1" {
		t.Log(todoTasks[0].Key)
		t.Errorf("Failed")
	}
}

func TestBuildKanban(t *testing.T) {

	var folderPath = ".././unittest/myTasks"

	var kanban Kanban = BuildKanban(folderPath)

	if len(kanban.bans) != 4 || kanban.bans[0].Name != "Todo" {
		t.Log(kanban)
		t.Errorf("Failed")
	}
}
