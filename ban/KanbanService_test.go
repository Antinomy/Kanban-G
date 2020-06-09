package ban

import (
	kt "kanban/task"
	"testing"
)

func TestKanSpec(t *testing.T) {

	var folderPath = ".././unittest/myTasks"

	var kanban Kanban = BuildKanban(folderPath)

	var kanSpec KanSpec = getKanSpec(kanban, kt.UNKNOWN)

	if len(kanSpec.hearders) != 4 {
		t.Log(kanSpec.hearders)
		t.Errorf("Failed")
	}

	if kanSpec.maxCellSize != 33 {
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

	var kanban Kanban = BuildKanban(folderPath)

	var kanSpec KanSpec = getKanSpec(kanban, kt.OWNER)

	if len(kanSpec.hearders) != 4 {
		t.Log(kanSpec.hearders)
		t.Errorf("Failed")
	}

	if kanSpec.maxCellSize != 33 {
		t.Log(kanSpec.maxCellSize)
		t.Errorf("Failed")
	}

	if len(kanSpec.rows) != 5 {
		t.Log(kanSpec.rows)
		t.Errorf("Failed")
	}
}

func TestKanSpecMap(t *testing.T) {

	var folderPath = ".././unittest/myTasks"

	var kanban Kanban = BuildKanban(folderPath)

	var kanSpec KanSpec = getKanSpec(kanban, kt.OWNER)

	var taskMap map[string]kt.Task = kanSpec.taskMap
	var banMap map[string]Ban = kanSpec.banMap

	if taskMap["t1"].FullName != "AY-H-ProjectA-0531-doSth.md" {
		t.Log(taskMap["t1"])
		t.Errorf("Failed")
		t.Log(taskMap)
	}

	if banMap["t1"].folder != "01-Todo" {
		t.Log(banMap["t1"])
		t.Errorf("Failed")
		t.Log(banMap)
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

func TestDeadlineType(t *testing.T) {
	var dl string

	dl = getDeadlineType("2020")
	if dl != kt.YEAR {
		t.Log(dl)
		t.Errorf("Failed")
	}

	dl = getDeadlineType("0512")
	if dl != "MONTH:05" {
		t.Log(dl)
		t.Errorf("Failed")
	}
}

func TestGetBan(t *testing.T) {
	var folderPath = ".././unittest/myTasks"
	var kanban Kanban = BuildKanban(folderPath)
	var ban Ban = getBan(kanban, "i")

	if ban.folder != "02-Doing" {
		t.Log(ban)
		t.Errorf("Failed")
	}

	ban = getBan(kanban, "ii")

	if ban.folder == "nil" {
		t.Log(ban.folder)
		t.Errorf("Failed")
	}

}

func TestKan(t *testing.T) {

	var folderPath = ".././unittest/myTasks"
	// folderPath = "/Users/Antinomy/Github/MyTask"

	var kanban Kanban = BuildKanban(folderPath)

	Kan(kanban, kt.UNKNOWN)
	Kan(kanban, kt.OWNER)
	Kan(kanban, kt.PRIORITY)
	Kan(kanban, kt.PROJECT)
	Kan(kanban, kt.DEADLINE)
}
