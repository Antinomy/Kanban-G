package ban

import (
	kc "kanban/conf"
	kt "kanban/task"
)

func buildKanban(folderPath string) Kanban {

	var configPath = ".././conf/conf.json"
	var config kc.Jconf = readJsonConfig(configPath)

	var banconfigs []kc.BanConfig = config.BanConfigs

	var result Kanban
	var bans []Ban

	for _, banConf := range banconfigs {

		var ban Ban
		var fullBanPath string = folderPath + "/" + banConf.Folder
		ban.folder = banConf.Folder

		var taskNameList []string = readFileList(fullBanPath)
		var tasks []kt.Task = readCorrectTasks(taskNameList)
		ban.tasks = tasks
		bans = append(bans, ban)
	}

	result.rootPath = folderPath
	result.bans = bans
	return result
}
