package ban

import (
	"fmt"
	kt "kanban/task"
	"sort"
	"strings"
	"time"

	"github.com/bndr/gotabulate"
)

func Kan(kanban Kanban, taskItem kt.TaskItem) {

	var kanSpec KanSpec = getKanSpec(kanban, taskItem)

	var ts kt.TaskService = new(kt.FileWay)

	tskItem := string(taskItem)

	var viewType = ts.FillBlank("Kan View : "+tskItem, kanSpec.maxCellSize)

	var today = ts.FillBlank("Today    : "+time.Now().Format("0102"), kanSpec.maxCellSize)

	println(viewType)
	println(today)
	if kanban.IsShortMode {
		println("ShortMode:", kanban.IsShortMode)
	}

	// Create an object from 2D interface array

	table := gotabulate.Create(kanSpec.rows)

	table.SetHeaders(kanSpec.hearders)

	// Set the Empty String (optional)
	// t.SetEmptyString("")

	// Set Align (Optional)
	table.SetAlign("left")

	// Set Max Cell Size
	table.SetMaxCellSize(kanSpec.maxCellSize)

	// Turn On String Wrapping
	table.SetWrapStrings(true)

	// Print the result: grid, or simple
	fmt.Println(table.Render("grid"))

}

func getKanSpec(kanban Kanban, taskItem kt.TaskItem) KanSpec {

	var result KanSpec
	var ts kt.TaskService = new(kt.FileWay)

	taskMap := make(map[string]kt.Task)
	banMap := make(map[string]Ban)

	// calc headers and maxCellSize
	calcInfo(&result, &kanban, ts, taskItem)

	if taskItem == kt.UNKNOWN {
		var cols []interface{}

		for _, ban := range kanban.bans {

			if kanban.IsShortMode && ban.supportShortMode {
				continue
			}

			var cell string
			for _, tk := range ban.tasks {
				var taskDesc = ts.GetTaskDesc(tk, taskItem, kanban.IsShortMode)
				taskDesc = ts.FillBlank(taskDesc, result.maxCellSize)
				cell += taskDesc
				banMap[tk.Key] = ban
				taskMap[tk.Key] = tk
			}
			cols = append(cols, cell)

		}
		result.rows = append(result.rows, cols)

		result.taskMap = taskMap
		result.banMap = banMap
		return result
	}

	var rowLines []string

	if taskItem == kt.OWNER {
		rowLines = result.owners
	}

	if taskItem == kt.PRIORITY {
		rowLines = result.priorities
	}

	if taskItem == kt.PROJECT {
		rowLines = result.projects
	}

	if taskItem == kt.DEADLINE {
		rowLines = result.deadlineTypes
	}

	for _, rowName := range rowLines {
		var cols []interface{}

		for index, ban := range kanban.bans {
			if index == 0 {
				cols = append(cols, rowName)
			}

			if kanban.IsShortMode && ban.supportShortMode {
				continue
			}

			var cell string
			for _, tk := range ban.tasks {
				var isOwnerCase bool = (rowName == tk.Owner && taskItem == kt.OWNER)
				var isPriorityCase bool = (rowName == tk.Priority && taskItem == kt.PRIORITY)
				var isProjectCase bool = (rowName == tk.Project && taskItem == kt.PROJECT)

				var isDeadlineCase bool = (rowName == getDeadlineType(tk.Deadline) && taskItem == kt.DEADLINE)

				if isOwnerCase || isPriorityCase || isProjectCase || isDeadlineCase {
					var taskDesc = ts.GetTaskDesc(tk, taskItem, kanban.IsShortMode)
					taskDesc = ts.FillBlank(taskDesc, result.maxCellSize)
					cell += taskDesc
				}

				banMap[tk.Key] = ban
				taskMap[tk.Key] = tk
			}
			cols = append(cols, cell)

		}
		result.rows = append(result.rows, cols)
	}

	result.taskMap = taskMap
	result.banMap = banMap
	return result
}

func calcInfo(result *KanSpec, kanban *Kanban, ts kt.TaskService, taskItem kt.TaskItem) {
	var maxCellSize int = 0

	result.priorities = appendPriorities(result.priorities)

	for _, ban := range kanban.bans {

		if kanban.IsShortMode && ban.supportShortMode {
			continue
		}

		var banHeader string = ban.Name + " [" + ban.prefix + "]"
		result.hearders = append(result.hearders, banHeader)

		for _, tk := range ban.tasks {

			result.owners = appendUnique(result.owners, tk.Owner)
			result.priorities = appendUnique(result.priorities, tk.Priority)
			result.projects = appendUnique(result.projects, tk.Project)
			result.deadlineTypes = appendUnique(result.deadlineTypes, getDeadlineType(tk.Deadline))

			var cellSize = len(ts.GetTaskDesc(tk, taskItem, kanban.IsShortMode))

			if maxCellSize < cellSize {
				maxCellSize = cellSize
			}

		}
	}

	sort.Strings(result.owners)
	sort.Strings(result.projects)
	sort.Strings(result.deadlineTypes)

	result.maxCellSize = maxCellSize
}

func appendUnique(result []string, target string) []string {
	for _, item := range result {
		if item == target {
			return result
		}
	}

	result = append(result, target)

	return result
}

func appendPriorities(result []string) []string {

	result = append(result, "H")
	result = append(result, "M")
	result = append(result, "L")

	return result
}

func getDeadlineType(deadline string) string {
	if strings.HasPrefix(deadline, "2") {
		return kt.YEAR
	}

	var today = time.Now().Format("0102")
	if deadline == today {
		return "DAY:" + today
	}

	return kt.MONTH + ":" + deadline[0:2]
}

func GetBan(kanban Kanban, banKey string) Ban {
	// set default Ban
	var result Ban = kanban.bans[0]

	for _, ban := range kanban.bans {
		if ban.prefix == banKey {
			result = ban
			return result
		}

		if strings.EqualFold(ban.Name, banKey) {
			result = ban
			return result
		}
	}

	return result
}
