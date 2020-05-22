package ban

import (
	kt "kanban/task"
	"testing"
)

func TestKanSpec(t *testing.T) {

	var folderPath = ".././unittest/myTasks"

	var kanban Kanban = buildKanban(folderPath)

	var kanSpec KanSpec = getKanSpec(kanban, kt.UNKNOWN)

	if len(kanSpec.hearders) != 4 {
		t.Log(kanSpec.hearders)
		t.Errorf("Failed")
	}

	if kanSpec.maxCellSize != 28 {
		t.Log(kanSpec.maxCellSize)
		t.Errorf("Failed")
	}

	if len(kanSpec.rows) != 1 {
		t.Log(len(kanSpec.rows))
		t.Errorf("Failed")
	}
}

func TestKan(t *testing.T) {

	var folderPath = ".././unittest/myTasks"
	// var folderPath = "/Users/Antinomy/Github/MyTask"

	var kanban Kanban = buildKanban(folderPath)

	kan(kanban, kt.UNKNOWN)
}
