package ban

import (
	"fmt"
	kc "kanban/conf"
	kt "kanban/task"
)

func readCorrectTasks(filesList []string) []kt.Task {
	var result []kt.Task

	var ts kt.TaskService = new(kt.FileWay)

	for _, fileName := range filesList {
		if ts.IsATask(fileName) {

			task := ts.CreateTask(fileName)
			fmt.Println("TaskCreated", task)
			result = append(result, task)
		}
	}

	fmt.Println("Validated Task Num", len(result))

	return result
}

func buildKanban(folderPath string) Kanban {

	var config kc.Jconf = loadConfig()

	var banconfigs []kc.BanConfig = config.BanConfigs

	var result Kanban
	var bans []Ban

	for _, banConf := range banconfigs {

		var ban Ban
		var fullBanPath string = folderPath + "/" + banConf.Folder
		ban.folder = banConf.Folder
		ban.name = banConf.Name

		var taskNameList []string = readFileList(fullBanPath)
		var tasks []kt.Task = readCorrectTasks(taskNameList)
		ban.tasks = tasks
		bans = append(bans, ban)
	}

	result.rootPath = folderPath
	result.bans = bans
	return result
}

func loadConfig() kc.Jconf {
	var configPath = ".././conf/conf.json"
	var config kc.Jconf = readJsonConfig(configPath)

	return config
}
