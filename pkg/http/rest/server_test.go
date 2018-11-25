package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/danillouz/todo"
	"github.com/danillouz/todo/pkg/http/rest"

	"github.com/danillouz/todo/mock"
)

func TestHandleTodos(t *testing.T) {
	service := &mock.TaskService{}
	server := rest.NewServer(service)

	t.Run("GET /todos", func(t *testing.T) {
		tasks := todo.Tasks{
			todo.Task{
				ID:    1,
				Descr: "Buy milk",
				Done:  false,
			},
		}

		service.TasksFn = func() (todo.Tasks, error) {
			return tasks, nil
		}

		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/todos", nil)

		server.ServeHTTP(w, r)

		t.Run("Status 200", func(t *testing.T) {
			got := w.Code
			want := http.StatusOK

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run("Content-Type application/json", func(t *testing.T) {
			got := w.Header().Get("content-type")
			want := "application/json"

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run("lists tasks as JSON", func(t *testing.T) {
			var got todo.Tasks
			json.NewDecoder(w.Body).Decode(&got)

			if !reflect.DeepEqual(got, tasks) {
				t.Errorf("got %v, want %v", got, tasks)
			}
		})
	})

	t.Run("POST /todos", func(t *testing.T) {
		task := todo.Task{
			ID:    0,
			Descr: "Buy eggs",
			Done:  false,
		}

		service.AddTaskFn = func(descr string) (todo.Task, error) {
			return task, nil
		}

		w := httptest.NewRecorder()
		d := []byte(`{"description": "Buy eggs"}`)
		b := bytes.NewBuffer(d)
		r, _ := http.NewRequest(http.MethodPost, "/todos", b)

		server.ServeHTTP(w, r)

		t.Run("Status 201", func(t *testing.T) {
			got := w.Code
			want := http.StatusCreated

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run("Content-Type application/json", func(t *testing.T) {
			got := w.Header().Get("content-type")
			want := "application/json"

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run("creates task and returns as JSON", func(t *testing.T) {
			var got todo.Task
			json.NewDecoder(w.Body).Decode(&got)

			if !reflect.DeepEqual(got, task) {
				t.Errorf("got %v, want %v", got, task)
			}
		})
	})
}
