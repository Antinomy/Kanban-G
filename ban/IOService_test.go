package ban

import (
	kc "kanban/conf"
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {

	var folderPath = ".././unittest/myTasks/02-doing"
	var doingList []string = readFileList(folderPath)

	var taskNum = len(doingList)

	if taskNum != 4 {
		t.Log(doingList)
		t.Errorf("Failed")
	}
}

func TestReadConfig(t *testing.T) {
	var config kc.Jconf = loadConfig()

	var banconfigs []kc.BanConfig = config.BanConfigs

	if len(banconfigs) != config.BanSize {
		t.Log(banconfigs)
		t.Errorf("Failed")
	}

	if banconfigs[0].Name != "Todo" {
		t.Log(banconfigs[0].Name)
		t.Errorf("Failed")
	}

	if banconfigs[0].Folder != "01-Todo" {
		t.Log(banconfigs[0].Folder)
		t.Errorf("Failed")
	}

	if banconfigs[0].Prefix != "t" {
		t.Log(banconfigs[0].Prefix)
		t.Errorf("Failed")
	}

	if banconfigs[3].SupportShortMode != true {
		t.Log(banconfigs[3].SupportShortMode)
		t.Errorf("Failed")
	}
}

func TestMoveFile(t *testing.T) {

	var folderPath = ".././unittest/myTasks"

	var existFile = "03-Hold/ZZ-H-ProjectZ-2020-doSth.md"

	var renameFile = "03-Hold/AA-H-ProjectZ-2020-doSth.md"

	var err = moveFile(folderPath, existFile, renameFile)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

	err = moveFile(folderPath, renameFile, existFile)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

	err = moveFile(folderPath, renameFile+"9999", existFile)

	if err == nil {
		t.Log(err)
		t.Errorf("Failed")
	}

}

func TestCreateFile(t *testing.T) {

	var folderPath = ".././unittest/myTasks"

	var newFile = "ZZ-H-ProjectZ-9999-doSth.md"

	var fullFilePath = folderPath + "/01-Todo/" + newFile
	var err = createFile(fullFilePath)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

	err = os.Remove(fullFilePath)

	if err != nil {
		t.Log(err)
		t.Errorf("Failed")
	}

}
