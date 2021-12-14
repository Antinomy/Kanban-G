package ban

import (
	kc "kanban/conf"
	kt "kanban/task"
	"strconv"
)

func readCorrectTasks(filesList []string, key string) []kt.Task {
	var result []kt.Task

	var ts kt.TaskService = new(kt.FileWay)

	var index int = 1
	for _, fileName := range filesList {
		if ts.IsATask(fileName) {

			task := ts.CreateTask(fileName)
			task.Key = key + strconv.Itoa(index)
			// fmt.Println("TaskCreated", task)
			result = append(result, task)
			index++
		}
	}

	// fmt.Println("Validated Task Num", len(result))

	return result
}

func BuildKanban(folderPath string) Kanban {

	var config kc.Jconf = loadConfig(folderPath)

	var banconfigs []kc.BanConfig = config.BanConfigs

	var result Kanban
	var bans []Ban

	for _, banConf := range banconfigs {

		var ban Ban
		var fullBanPath string = folderPath + "/" + banConf.Folder
		ban.folder = banConf.Folder
		ban.Name = banConf.Name
		ban.supportShortMode = banConf.SupportShortMode
		ban.prefix = banConf.Prefix

		var taskNameList []string = readFileList(fullBanPath)
		var tasks []kt.Task = readCorrectTasks(taskNameList, ban.prefix)
		ban.tasks = tasks
		bans = append(bans, ban)
	}

	result.rootPath = folderPath
	result.bans = bans
	return result
}
