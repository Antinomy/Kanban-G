package ban

import (
	kt "kanban/task"
	"testing"
)

func TestChangeTask(t *testing.T) {

	var folderPath = ".././unittest/myTasks"

	var kanban Kanban = BuildKanban(folderPath)

	var kanSpec KanSpec = getKanSpec(kanban, kt.OWNER)

	var changeSpec ChangeSpec = ChangeTask(kanSpec, "t1", "o", "ZZ")

	if changeSpec.originPath != "01-Todo/AY-H-ProjectA-0531-doSth.md" {
		t.Log(changeSpec)
		t.Errorf("Failed")
	}

	if changeSpec.changedPath != "01-Todo/ZZ-H-ProjectA-0531-doSth.md" {
		t.Log(changeSpec)
		t.Errorf("Failed")
	}

}

func TestChangeBan(t *testing.T) {

	var folderPath = ".././unittest/myTasks"

	var kanban Kanban = BuildKanban(folderPath)

	var kanSpec KanSpec = getKanSpec(kanban, kt.OWNER)

	var changeSpec ChangeSpec = ChangeBan(kanban, kanSpec, "t1", "i")

	if changeSpec.originPath != "01-Todo/AY-H-ProjectA-0531-doSth.md" {
		t.Log(changeSpec)
		t.Errorf("Failed")
	}

	if changeSpec.changedPath != "02-Doing/AY-H-ProjectA-0531-doSth.md" {
		t.Log(changeSpec)
		t.Errorf("Failed")
	}
}

func TestChange(t *testing.T) {

	var folderPath = ".././unittest/myTasks"
	var kanban Kanban = BuildKanban(folderPath)
	var kanSpec KanSpec = getKanSpec(kanban, kt.OWNER)

	var changeSpec ChangeSpec = ChangeBan(kanban, kanSpec, "h2", "i")

	var err = ChangeAny(folderPath, changeSpec)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

	kanban = BuildKanban(folderPath)
	kanSpec = getKanSpec(kanban, kt.OWNER)

	changeSpec = ChangeBan(kanban, kanSpec, "i4", "h")

	err = ChangeAny(folderPath, changeSpec)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

}
