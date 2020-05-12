package kanban

import (
    "testing"
)


func test(t *testing.T) {
    var Task task

    task.Owner = "AY"
    task.Priority = "M"
    task.Project = "ProjectA"
    task.Deadline = "20200512"
    task.Tittle = "WriteKanbanCode.md"

    t.Log(task)
}