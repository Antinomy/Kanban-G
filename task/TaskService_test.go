package task

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
    var taskService TaskService = new(FileWay)
    var task1 Task = taskService.createTask("AY-M-ProjectA-20200512-WriteKanbanCode.md")
    t.Log(task1)

    if task1.project != "ProjectA" {
        t.Errorf("Failed")
    }
}

func TestIsATask(t *testing.T) {
    var taskService TaskService = new(FileWay)
    var task1 bool = taskService.isATask("AY-M-ProjectA-20200512-WriteKanbanCode.md")

    if task1 == false {
        t.Errorf("Failed")
    }

    var task2 = taskService.isATask("whatever")
    if task2 == true {
        t.Errorf("Failed")
    }

}

func TestChangeTask(t *testing.T) {
    var taskService TaskService = new(FileWay)

    var originTask Task = taskService.createTask("AY-M-ProjectA-20200512-WriteKanbanCode.md")
    t.Log(originTask)

    if originTask.project != "ProjectA" {
        t.Errorf("Failed")
    }

    changingTask := ChangingTask{
        origin: originTask,
        changeItem: "Owner",
        changeContent: "WGL",
    }
    var changedTask Task = taskService.changeTask(changingTask);
    t.Log(changedTask)

    if changedTask.owner != "WGL" {
        t.Errorf("Failed")
    }
}
