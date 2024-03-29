package ban

import (
	"fmt"
	kt "kanban/task"
	"os"
	"strings"
	"testing"
)

func TestChangeTask(t *testing.T) {

	var folderPath = ".././unittest/myTasks"

	var kanban Kanban = BuildKanban(folderPath)

	var changeSpec ChangeSpec = ChangeTask(kanban, kt.OWNER, "t1", "o", "ZZ")

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

	var changeSpec ChangeSpec = ChangeBan(kanban, kt.OWNER, "t1", "i")

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

	// target file 03-Hold/ZZ-H-ProjectZ-2020-doSth.md
	// change ban to 02-Doing
	var changeSpec ChangeSpec = ChangeBan(kanban, kt.OWNER, "ho1", "i")

	var err = ChangeOne(folderPath, changeSpec)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

	kanban = BuildKanban(folderPath)

	// change ban back to 03-Hold
	changeSpec = ChangeBan(kanban, kt.OWNER, "i4", "ho")

	err = ChangeOne(folderPath, changeSpec)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

	kanban = BuildKanban(folderPath)

	// change task owner to AY
	changeSpec = ChangeTask(kanban, kt.OWNER, "ho1", "o", "AY")
	err = ChangeOne(folderPath, changeSpec)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

	kanban = BuildKanban(folderPath)

	// change task owner back to ZZ
	changeSpec = ChangeTask(kanban, kt.OWNER, "ho1", "o", "ZZ")
	err = ChangeOne(folderPath, changeSpec)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

}

func TestCreareBanTask(t *testing.T) {

	var folderPath = ".././unittest/myTasks"
	var newTask = "ZZ-H-ProjectZ-9999-doSth.md"

	var kanban Kanban = BuildKanban(folderPath)

	result, fullFilePath := CreateBanTask(kanban, newTask, "t")

	if result != true {
		t.Log(result)
		t.Errorf("Failed")
	}

	err := os.Remove(fullFilePath)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

}

func TestOpenTask(t *testing.T) {

	var folderPath = ".././unittest/myTasks"
	var newTask = "t1"

	var kanban Kanban = BuildKanban(folderPath)

	err := OpenTask(kanban, newTask, kt.OWNER)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

}

func TestExec(t *testing.T) {
	out, err := Exec(".././unittest/myTasks/countGit.sh")
	if err != nil {
		fmt.Println(err)
	}

	out = strings.Trim(out, "\n")
	println(out)

}
