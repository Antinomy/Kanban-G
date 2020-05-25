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

func TestKanSpecOwner(t *testing.T) {

	var folderPath = ".././unittest/myTasks"

	var kanban Kanban = buildKanban(folderPath)

	var kanSpec KanSpec = getKanSpec(kanban, kt.OWNER)

	if len(kanSpec.hearders) != 4 {
		t.Log(kanSpec.hearders)
		t.Errorf("Failed")
	}

	if kanSpec.maxCellSize != 28 {
		t.Log(kanSpec.maxCellSize)
		t.Errorf("Failed")
	}

	if len(kanSpec.rows) != 4 {
		t.Log(kanSpec.rows)
		t.Errorf("Failed")
	}
}

func TestUniqueArray(t *testing.T) {

	var owners []string
	owners = []string{"AY", "WL", "LL"}

	owners = appendUnique(owners, "xx")

	if len(owners) != 4 {
		t.Log(len(owners))
		t.Errorf("Failed")
	}

	owners = appendUnique(owners, "AY")

	if len(owners) != 4 {
		t.Log(len(owners))
		t.Errorf("Failed")
	}

}

func TestKan(t *testing.T) {

	var folderPath = ".././unittest/myTasks"
	// var folderPath = "/Users/Antinomy/Github/MyTask"

	var kanban Kanban = buildKanban(folderPath)

	kan(kanban, kt.UNKNOWN)
	kan(kanban, kt.OWNER)
	kan(kanban, kt.PRIORITY)
	kan(kanban, kt.PROJECT)
}
