package ban

import (
	kt "kanban/task"
)

// Ban entiy
type Ban struct {
	name   string
	folder string
	tasks  []kt.Task
}

// Kanban entiy
type Kanban struct {
	rootPath string
	bans     []Ban
}

// KanSpec entiy
type KanSpec struct {
	hearders    []string
	maxCellSize int
	rows        [][]interface{}
}
