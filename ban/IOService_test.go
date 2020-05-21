package ban

import (
	kc "kanban/conf"
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

	if banconfigs[0].Name != "todo" {
		t.Log(banconfigs[0].Name)
		t.Errorf("Failed")
	}

	if banconfigs[0].Folder != "01-todo" {
		t.Log(banconfigs[0].Folder)
		t.Errorf("Failed")
	}
}
