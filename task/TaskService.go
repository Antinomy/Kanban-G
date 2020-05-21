package task

import (
	"strings"
)

//TaskService interface e
type TaskService interface {
	CreateTask(taskName string) Task
	IsATask(taskName string) bool
	ChangeTask(changingTask ToChangeTask) Task
	GetTaskDesc(task Task, taskItem TaskItem) string
	FillBlank(taskDesc string, maxSize int) string
}

//FileWay desc
type FileWay struct {
}

func (t *FileWay) CreateTask(taskName string) Task {
	var arrs = strings.Split(taskName, "-")
	var result Task
	result.owner = arrs[0]
	result.priority = arrs[1]
	result.project = arrs[2]
	result.deadline = arrs[3]
	result.tittle = arrs[4]
	result.fullName = taskName

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
	var origin = changingTask.origin

	var result Task = origin

	if changingTask.changeItem == OWNER {
		result.owner = changingTask.changeContent
	}

	if changingTask.changeItem == PRIORITY {
		result.priority = changingTask.changeContent
	}

	if changingTask.changeItem == DEADLINE {
		result.deadline = changingTask.changeContent
	}

	return result
}

func (t *FileWay) GetTaskDesc(task Task, taskItem TaskItem) string {

	var result string = task.fullName

	if taskItem == UNKNOWN {
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
