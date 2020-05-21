package ban

import (
	"fmt"
	kt "kanban/task"

	"github.com/bndr/gotabulate"
)

func getKanSpec(kanban Kanban, taskItem kt.TaskItem) KanSpec {

	var result KanSpec
	var maxCellSize int = 0
	var ts kt.TaskService = new(kt.FileWay)

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

	if taskItem == kt.UNKNOWN {
		var cols []interface{}

		for _, ban := range kanban.bans {

			var cell string
			for _, tk := range ban.tasks {

				cell += ts.GetTaskDesc(tk, kt.UNKNOWN)
			}
			cols = append(cols, cell)

		}
		result.rows = append(result.rows, cols)

	}

	return result
}

func kan(kanban Kanban) {

	//clear screen
	print("\033[H\033[2J")

	var kanSpec KanSpec = getKanSpec(kanban, kt.UNKNOWN)

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
