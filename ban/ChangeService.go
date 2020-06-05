package ban

import (
	kt "kanban/task"
)

func ChangeTask(kanSpec KanSpec, taskKey string, taskItemStr string, changeContext string) ChangeSpec {
	var result ChangeSpec

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

func ChangeBan(kanban Kanban, kanSpec KanSpec, taskKey string, banStr string) ChangeSpec {
	var result ChangeSpec

	var originTask = kanSpec.taskMap[taskKey]
	var originBan = kanSpec.banMap[taskKey]
	var changedBan Ban = getBan(kanban, banStr)
	result.originPath = originBan.folder + "/" + originTask.FullName
	result.changedPath = changedBan.folder + "/" + originTask.FullName

	return result
}
