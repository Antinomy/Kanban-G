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

			task := ts.CreateTask(fileName)
			fmt.Println("TaskCreated", task)
			result = append(result, task)
		}
	}

	fmt.Println("Validated Task Num", len(result))

	return result
}
