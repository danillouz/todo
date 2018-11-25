package mock

import (
	"github.com/danillouz/todo"
)

// Ensure TaskService mock satisfies the TaskService interface.
var _ todo.TaskService = &TaskService{}

// TaskService mock.
type TaskService struct {
	AddTaskFn     func(descr string) (todo.Task, error)
	AddTaskCalled bool

	TasksFn     func() (todo.Tasks, error)
	TasksCalled bool

	UpdateTaskFn     func(id todo.TaskID, done bool) (todo.Task, error)
	UpdateTaskCalled bool

	DeleteTaskFn     func(id todo.TaskID) error
	DeleteTaskCalled bool
}

// AddTask mock.
func (s *TaskService) AddTask(descr string) (todo.Task, error) {
	s.AddTaskCalled = true

	return s.AddTaskFn(descr)
}

// Tasks mock.
func (s *TaskService) Tasks() (todo.Tasks, error) {
	s.TasksCalled = true

	return s.TasksFn()
}

// UpdateTask mock.
func (s *TaskService) UpdateTask(id todo.TaskID, done bool) (todo.Task, error) {
	s.UpdateTaskCalled = true

	return s.UpdateTaskFn(id, done)
}

// DeleteTask mock.
func (s *TaskService) DeleteTask(id todo.TaskID) error {
	s.DeleteTaskCalled = true

	return s.DeleteTaskFn(id)
}
