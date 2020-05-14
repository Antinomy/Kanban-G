package task



// Task entiy 
type Task struct {
    owner string
    priority string
    project string
    deadline string
    tittle string
 }

//ChangingTask desc
type ChangingTask struct{
    origin Task 
    changeItem string
    changeContent string
}