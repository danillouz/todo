package todo

// TaskID uniquely identifies a task.
type TaskID int

// Task represents something that needs to be done, i.e a "todo".
type Task struct {
	ID    TaskID `json:"id"`
	Descr string `json:"description"`
	Done  bool   `json:"done"`
}

// Tasks represents a list of tasks.
type Tasks []Task

// TaskService represents a service to manage tasks.
type TaskService interface {
	AddTask(descr string) (Task, error)
	Tasks() (Tasks, error)
	UpdateTask(id TaskID, done bool) (Task, error)
	DeleteTask(id TaskID) error
}
