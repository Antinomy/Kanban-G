package ban

import (
	"fmt"
	kt "kanban/task"

	"github.com/bndr/gotabulate"
)

func kan(kanban Kanban, taskItem kt.TaskItem) {

	//clear screen
	print("\033[H\033[2J")

	var kanSpec KanSpec = getKanSpec(kanban, taskItem)

	// Create an object from 2D interface array

	t := gotabulate.Create(kanSpec.rows)

	t.SetHeaders(kanSpec.hearders)

	// Set the Empty String (optional)
	// t.SetEmptyString("")

	// Set Align (Optional)
	t.SetAlign("left")

	// Set Max Cell Size
	t.SetMaxCellSize(kanSpec.maxCellSize)

	// Turn On String Wrapping
	t.SetWrapStrings(true)

	// Print the result: grid, or simple
	fmt.Println(t.Render("grid"))

}

func getKanSpec(kanban Kanban, taskItem kt.TaskItem) KanSpec {

	var result KanSpec
	var ts kt.TaskService = new(kt.FileWay)

	// calc headers and maxCellSize
	calcHeaderAndMaxCellSize(&result, &kanban, ts, taskItem)

	if taskItem == kt.UNKNOWN {
		var cols []interface{}

		for _, ban := range kanban.bans {

			var cell string
			for _, tk := range ban.tasks {
				var taskDesc = ts.GetTaskDesc(tk, taskItem)
				taskDesc = ts.FillBlank(taskDesc, result.maxCellSize)
				cell += taskDesc
			}
			cols = append(cols, cell)

		}
		result.rows = append(result.rows, cols)
	}

	return result
}

func calcHeaderAndMaxCellSize(result *KanSpec, kanban *Kanban, ts kt.TaskService, taskItem kt.TaskItem) {
	var maxCellSize int = 0

	for _, ban := range kanban.bans {
		result.hearders = append(result.hearders, ban.name)

		for _, tk := range ban.tasks {

			var cellSize = len(ts.GetTaskDesc(tk, taskItem))

			if maxCellSize < cellSize {
				maxCellSize = cellSize
			}

		}
	}
	result.maxCellSize = maxCellSize
}
