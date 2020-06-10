package ban

import (
	"fmt"
	kt "kanban/task"
	"log"
)

func ChangeTask(kanban Kanban, existTaskItem kt.TaskItem, taskKey string, taskItemStr string, changeContext string) ChangeSpec {
	var result ChangeSpec

	var kanSpec KanSpec = getKanSpec(kanban, existTaskItem)

	var taskItem kt.TaskItem = kt.GetTaskItem(taskItemStr)
	var originTask = kanSpec.taskMap[taskKey]

	changingTask := kt.ToChangeTask{
		Origin:        originTask,
		ChangeItem:    taskItem,
		ChangeContent: changeContext,
	}

	var taskService kt.TaskService = new(kt.FileWay)
	var changedTask kt.Task = taskService.ChangeTask(changingTask)

	folderPath := kanSpec.banMap[taskKey].folder + "/"
	result.originPath = folderPath + originTask.FullName
	result.changedPath = folderPath + changedTask.FullName

	return result
}

func ChangeBan(kanban Kanban, taskItem kt.TaskItem, taskKey string, prefix string) ChangeSpec {
	var result ChangeSpec

	var kanSpec KanSpec = getKanSpec(kanban, taskItem)
	var originTask = kanSpec.taskMap[taskKey]
	var originBan = kanSpec.banMap[taskKey]
	var changedBan Ban = getBan(kanban, prefix)
	result.originPath = originBan.folder + "/" + originTask.FullName
	result.changedPath = changedBan.folder + "/" + originTask.FullName

	return result
}

func ChangeOne(folderPath string, changeSpec ChangeSpec) error {
	var result error = moveFile(folderPath, changeSpec.originPath, changeSpec.changedPath)
	return result
}

func CreateBanTask(kanban Kanban, newTask string, prefix string) (bool, string) {

	var taskService kt.TaskService = new(kt.FileWay)
	var result bool = taskService.IsATask(newTask)

	if !result {
		fmt.Println("Task Format Error: ", newTask, ",  e.g ZZ-H-ProjectZ-9999-doSth.md")
		return result, "NoPath"
	}

	var createInBan Ban = getBan(kanban, prefix)
	var fullFilePath = kanban.rootPath + "/" + createInBan.folder + "/" + newTask

	var err = createFile(fullFilePath)

	if err != nil {
		log.Println(err)
		result = false
	}

	return result, fullFilePath
}
