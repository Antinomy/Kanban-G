package task

import (
	"strings"
)

//TaskService interface e
type TaskService interface {
    createTask(taskName string) Task
    isATask(taskName string) bool
    changeTask(changingTask ChangingTask) Task
}

 //FileWay desc
 type FileWay struct{
}

func (t *FileWay)createTask(taskName string) Task {
    var arrs = strings.Split(taskName, "-")
    var result Task
    result.owner = arrs[0]
    result.priority = arrs[1]
    result.project = arrs[2]
    result.deadline = arrs[3]
    result.tittle = arrs[4]

    return result
}

func (t *FileWay) isATask(taskName string) bool {
    var result bool = false

    var arrs = strings.Split(taskName, "-")

    if len(arrs) == 5 {
        result = true
    }

    return result
}

func (t *FileWay) changeTask(changingTask ChangingTask) Task {
    var origin = changingTask.origin;

    var result Task =origin;

    if(changingTask.changeItem == "Owner"){
        result.owner = changingTask.changeContent
    }
   
    return result
}