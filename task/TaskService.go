package task

import (
	"strings"
)

//TaskService interface e
type TaskService interface {
	CreateTask(taskName string) Task
	IsATask(taskName string) bool
	ChangeTask(changingTask ToChangeTask) Task
	GetTaskDesc(task Task, taskItem TaskItem, isShortMode bool) string
	FillBlank(taskDesc string, maxSize int) string
}

//FileWay desc
type FileWay struct {
}

func (t *FileWay) CreateTask(taskName string) Task {
	var arrs = strings.Split(taskName, "-")
	var result Task
	result.Owner = arrs[0]
	result.Priority = arrs[1]
	result.Project = arrs[2]
	result.Deadline = arrs[3]
	result.Tittle = arrs[4]
	result.FullName = taskName

	return result
}

func (t *FileWay) IsATask(taskName string) bool {
	var result bool = false

	var arrs = strings.Split(taskName, "-")

	if len(arrs) == 5 {
		result = true
	}

	return result
}

func (t *FileWay) ChangeTask(changingTask ToChangeTask) Task {
	var origin = changingTask.Origin

	var result Task = origin

	if changingTask.ChangeItem == OWNER {
		result.Owner = changingTask.ChangeContent
	}

	if changingTask.ChangeItem == PRIORITY {
		result.Priority = changingTask.ChangeContent
	}

	if changingTask.ChangeItem == PROJECT {
		result.Project = changingTask.ChangeContent
	}

	if changingTask.ChangeItem == DEADLINE {
		result.Deadline = changingTask.ChangeContent
	}

	if changingTask.ChangeItem == TITTLE {
		result.Tittle = changingTask.ChangeContent
	}

	result.FullName = result.Owner + "-" + result.Priority + "-" + result.Project + "-" + result.Deadline + "-" + result.Tittle

	return result
}

func (t *FileWay) GetTaskDesc(task Task, taskItem TaskItem, isShortMode bool) string {

	var result string = task.FullName

	var prefix string = ""

	if len(task.Key) != 0 {
		prefix = "[" + task.Key + "] "
	}
	result = prefix + result

	if !isShortMode {
		return result
	}

	if taskItem == UNKNOWN {
		return result
	}

	if taskItem == OWNER {
		result = prefix + task.Priority + "-" + task.Project + "-" + task.Deadline + "-" + task.Tittle
		return result
	}

	if taskItem == PRIORITY {
		result = prefix + task.Owner + "-" + task.Project + "-" + task.Deadline + "-" + task.Tittle
		return result
	}

	if taskItem == PROJECT {
		result = prefix + task.Owner + "-" + task.Priority + "-" + task.Deadline + "-" + task.Tittle
		return result
	}

	if taskItem == DEADLINE {
		var dateLen = len(task.Deadline)
		shortDate := task.Deadline

		if dateLen > 2 {
			shortDate = shortDate[dateLen-2 : dateLen]
		}

		result = prefix + task.Owner + "-" + task.Priority + "-" + task.Project + "-" + shortDate + "-" + task.Tittle
		return result
	}

	return result
}

func (t *FileWay) FillBlank(taskDesc string, maxSize int) string {

	var currentSize = len(taskDesc)

	if currentSize >= maxSize {
		return taskDesc
	}

	var result string = taskDesc

	var fillSize = maxSize - currentSize

	for i := 0; i < fillSize; i++ {
		result += " "
	}

	return result
}
