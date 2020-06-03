package ban

import (
	kt "kanban/task"
)

// Ban entiy
type Ban struct {
	name             string
	folder           string
	tasks            []kt.Task
	supportShortMode bool
	prefix           string
}

// Kanban entiy
type Kanban struct {
	rootPath    string
	bans        []Ban
	IsShortMode bool
}

// KanSpec entiy
type KanSpec struct {
	hearders      []string
	maxCellSize   int
	owners        []string
	priorities    []string
	projects      []string
	deadlineTypes []string
	rows          [][]interface{}
}
