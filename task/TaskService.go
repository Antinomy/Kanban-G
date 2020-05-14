package task

import (
	"strings"
)

//taskService interface e
type taskService interface {
	toTask() Task
	isATask() bool
}

func toTask(fileName string) (result Task) {
	var arrs = strings.Split(fileName, "-")

	result.owner = arrs[0]
	result.priority = arrs[1]
	result.project = arrs[2]
	result.deadline = arrs[3]
	result.tittle = arrs[4]

	return result
}

func isATask(fileName string) (result bool) {
	result = false

	var arrs = strings.Split(fileName, "-")

	if len(arrs) == 5 {
		result = true
	}

	return result
}
