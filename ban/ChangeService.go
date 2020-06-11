package ban

import (
	"fmt"
	kt "kanban/task"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	var changedBan Ban = GetBan(kanban, prefix)
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

	var createInBan Ban = GetBan(kanban, prefix)
	var fullFilePath = kanban.rootPath + "/" + createInBan.folder + "/" + newTask

	var err = createFile(fullFilePath)

	if err != nil {
		log.Println(err)
		result = false
	}

	return result, fullFilePath
}

func OpenTask(kanban Kanban, taskKey string, taskItem kt.TaskItem) error {
	var kanSpec KanSpec = getKanSpec(kanban, taskItem)
	var originTask = kanSpec.taskMap[taskKey]
	var originBan = kanSpec.banMap[taskKey]

	var fullFilePath = kanban.rootPath + "/" + originBan.folder + "/" + originTask.FullName

	absPath, _ := filepath.Abs(fullFilePath)

	fmt.Println(absPath)

	// var vscodeCmd = "/usr/bin/open -n -b \"com.microsoft.VSCode\" --args "

	cmd := exec.Command("open", "-n", "-b", "com.microsoft.VSCode", "--args", absPath)
	cmd.Env = os.Environ()

	// fmt.Println(cmd.Env)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Execute Command ERROR : " + err.Error())
	}

	fmt.Println("Execute Command End.")
	return err
}
