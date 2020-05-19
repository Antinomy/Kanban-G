package ban

import (
	"fmt"
	kt "kanban/task"
)

func readCorrectTasks(filesList []string) []kt.Task {
	var result []kt.Task

	var ts kt.TaskService = new(kt.FileWay)

	for _, fileName := range filesList {
		if ts.IsATask(fileName) {
			fmt.Println(fileName)
			result = append(result, ts.CreateTask(fileName))
		}
	}

	fmt.Println(len(result))

	return result
}
