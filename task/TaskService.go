
package task

import (
    "strings"
)

// interface task service
type TaskService interface{
    toTask() Task
}


// func toTask (i TaskService){
//     i.toTask()
// }

func toTask(fileName string)(result Task){
  
    var arrs = strings.Split(fileName,"-")

    result.owner = arrs[0]
    result.priority = arrs[1]
    result.project = arrs[2]
    result.deadline = arrs[3]
    result.tittle = arrs[4]


	return result;
} 