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

// KanSpec vo
type KanSpec struct {
	hearders      []string
	maxCellSize   int
	owners        []string
	priorities    []string
	projects      []string
	deadlineTypes []string
	rows          [][]interface{}
	taskMap       map[string]kt.Task
	banMap        map[string]Ban
}

// ChangeSpec vo
type ChangeSpec struct {
	originPath  string
	changedPath string
}
