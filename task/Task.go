package task

import "strings"

// Task entiy
type Task struct {
	FullName string
	Owner    string
	Priority string
	Project  string
	Deadline string
	Tittle   string
	Key      string
}

// ToChangeTask desc
type ToChangeTask struct {
	Origin        Task
	ChangeItem    TaskItem
	ChangeContent string
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

	case "TITTLE", "TIT", "T":
		return TITTLE

	default:
		return UNKNOWN
	}
}

// deadline type desc
const (
	MONTH string = "MONTH"
	YEAR  string = "YEAR"
)
