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

// Ban entiy
type Kanban struct {
	rootPath string
	bans     []Ban
}

// Ban entiy
type KanSpec struct {
	hearders    []string
	maxCellSize int
	rows        [][]interface{}
}
