package task

import "strings"

// Task entiy
type Task struct {
	fullName string
	Owner    string
	Priority string
	Project  string
	Deadline string
	tittle   string
	Key      string
}

// ToChangeTask desc
type ToChangeTask struct {
	origin        Task
	changeItem    TaskItem
	changeContent string
}

// TaskItem desc
type TaskItem string

const (
	OWNER    TaskItem = "OWNER"
	PRIORITY TaskItem = "PRIORITY"
	DEADLINE TaskItem = "DEADLINE"
	TITTLE   TaskItem = "TITTLE"
	PROJECT  TaskItem = "PROJECT"
	UNKNOWN  TaskItem = "UNKNOWN"
)

func GetTaskItem(itemTypeStr string) TaskItem {
	var itemType string = strings.ToUpper(itemTypeStr)

	switch itemType {
	case "OWNER", "OWN", "O":
		return OWNER

	case "PRIORITY", "PRI", "PI", "I":
		return PRIORITY

	case "DEADLINE", "DL", "D":
		return DEADLINE

	case "PROJECT", "PRJ", "PJ", "J":
		return PROJECT

	default:
		return UNKNOWN
	}
}

// deadline type desc
const (
	MONTH string = "MONTH"
	YEAR  string = "YEAR"
)
