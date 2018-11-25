package memory

import "github.com/danillouz/todo"

// Ensure Storage satisfies the TaskService interface.
var _ todo.TaskService = &Storage{}

// Storage implements an in memory storage to store and manage tasks.
type Storage struct {
	todos map[todo.TaskID]todo.Task

	// FIXME
	id int
}

// NewStorage returns an empty storage.
func NewStorage() *Storage {
	todos := make(map[todo.TaskID]todo.Task)

	return &Storage{
		todos: todos,
	}
}

// AddTask adds a task to the in memory storage.
func (s *Storage) AddTask(descr string) (todo.Task, error) {
	id := todo.TaskID(s.id)
	task := todo.Task{
		ID:    id,
		Descr: descr,
	}

	s.todos[id] = task

	// FIXME
	s.id++

	return task, nil
}

// Tasks retrieves all tasks from the in memory storage.
func (s *Storage) Tasks() (todo.Tasks, error) {
	tasks := todo.Tasks{}

	for _, todo := range s.todos {
		tasks = append(tasks, todo)
	}

	return tasks, nil
}

// UpdateTask updates a single task in memory storage.
func (s *Storage) UpdateTask(id todo.TaskID, done bool) (todo.Task, error) {
	task := todo.Task{}

	return task, nil
}

// DeleteTask deletes a task from in memory storage.
func (s *Storage) DeleteTask(id todo.TaskID) error {
	return nil
}
