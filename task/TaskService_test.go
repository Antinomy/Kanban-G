package task

import (
    "testing"
)


func TestTaskService(t *testing.T) {
    // var  taskSrv TaskService

    var fileName string = "AY-M-ProjectA-20200512-WriteKanbanCode.md"
	
	var task1 Task;
	task1 = toTask(fileName)

    t.Log(task1)

 
    if task1.project != "ProjectA" {
        t.Errorf("Failed")
    }
 
  
}