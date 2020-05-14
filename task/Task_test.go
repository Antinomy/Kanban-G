package task

import (
	"testing"
)


func TestTask(t *testing.T) {
    var  task1 Task

    task1.owner = "AY"
    task1.priority = "M"
    task1.project = "ProjectA"
    task1.deadline = "20200512"
    task1.tittle = "WriteKanbanCode.md"

    t.Log(task1)

 
    if task1.owner != "AY" {
        t.Errorf("Failed")
    }
 
  
}