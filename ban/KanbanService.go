package ban

import (
	"fmt"
	kt "kanban/task"

	"github.com/bndr/gotabulate"
)

func kan(kanban Kanban) {

	var result [][]interface{}

	//clear screen
	print("\033[H\033[2J")

	var cols []interface{}

	var ts kt.TaskService = new(kt.FileWay)

	for _, ban := range kanban.bans {

		var cell string
		for _, tk := range ban.tasks {

			cell += ts.GetTaskDesc(tk, kt.UNKNOWN)
		}
		cols = append(cols, cell)

	}
	result = append(result, cols)

	// Create an object from 2D interface array

	t := gotabulate.Create(result)

	// Set the Headers (optional)
	var hearders []string
	for _, ban := range kanban.bans {
		hearders = append(hearders, ban.name)
	}

	t.SetHeaders(hearders)

	// Set the Empty String (optional)
	// t.SetEmptyString("")

	// Set Align (Optional)
	t.SetAlign("left")

	// Set Max Cell Size
	t.SetMaxCellSize(40)

	// Turn On String Wrapping
	// t.SetWrapStrings(true)

	// Print the result: grid, or simple
	fmt.Println(t.Render("grid"))

}
