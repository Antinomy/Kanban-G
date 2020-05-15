package task

import "strings"

// Task entiy
type Task struct {
	owner    string
	priority string
	project  string
	deadline string
	tittle   string
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

func getTaskItem(itemTypeStr string) TaskItem {
	var itemType string = strings.ToUpper(itemTypeStr)

	switch itemType {
	case "OWNER", "OWN", "O":
		return OWNER

	case "PRIORITY", "POR", "PI":
		return PRIORITY

	case "DEADLINE", "DL", "D":
		return DEADLINE

	case "PROJECT", "PRJ", "PJ":
		return PROJECT

	default:
		return UNKNOWN
	}
}
