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

	// target file 03-Hold/ZZ-H-ProjectZ-2020-doSth.md
	// change ban to 02-Doing
	var changeSpec ChangeSpec = ChangeBan(kanban, kanSpec, "h2", "i")

	var err = ChangeOne(folderPath, changeSpec)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

	kanban = BuildKanban(folderPath)
	kanSpec = getKanSpec(kanban, kt.OWNER)

	// change ban back to 03-Hold
	changeSpec = ChangeBan(kanban, kanSpec, "i4", "h")

	err = ChangeOne(folderPath, changeSpec)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

	kanban = BuildKanban(folderPath)
	kanSpec = getKanSpec(kanban, kt.OWNER)

	// change task owner to AY
	changeSpec = ChangeTask(kanSpec, "h2", "o", "AY")
	err = ChangeOne(folderPath, changeSpec)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

	kanban = BuildKanban(folderPath)
	kanSpec = getKanSpec(kanban, kt.OWNER)

	// change task owner back to ZZ
	changeSpec = ChangeTask(kanSpec, "h2", "o", "ZZ")
	err = ChangeOne(folderPath, changeSpec)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

}
