package task

// Task entiy
type Task struct {
	owner    string
	priority string
	project  string
	deadline string
	tittle   string
}

//ChangingTask desc
type ChangingTask struct {
	origin        Task
	changeItem    TaskItem
	changeContent string
}

// TaskItem
type TaskItem string

const (
	OWNER    TaskItem = "OWNER"
	PRIORITY TaskItem = "PRIORITY"
	DEADLINE TaskItem = "DEADLINE"
	TITTLE   TaskItem = "TITTLE"
)
