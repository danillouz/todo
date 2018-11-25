package rest

import (
	"encoding/json"
	"net/http"

	"github.com/danillouz/todo"
)

// Server represents an HTTP server.
type Server struct {
	service todo.TaskService
	http.Handler
}

// NewServer creates a new HTTP server.
func NewServer(s todo.TaskService) *Server {
	server := new(Server)
	server.service = s

	router := http.NewServeMux()
	router.HandleFunc("/todos", server.handleTodos)
	server.Handler = router

	return server
}

// Run runs the HTTP server.
func (s Server) Run(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s Server) handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleCreateTodos(s, w, r)
	case http.MethodGet:
		handleReadTodos(s, w, r)
	}
}

func handleCreateTodos(s Server, w http.ResponseWriter, r *http.Request) {
	var body todo.Task
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	task, err := s.service.AddTask(body.Descr)

	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	sendJSON(w, task)
}

func handleReadTodos(s Server, w http.ResponseWriter, r *http.Request) {
	tasks, err := s.service.Tasks()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSON(w, tasks)
}

type errResponse struct {
	Err string `json:"error"`
}

func sendError(w http.ResponseWriter, err error, code int) {
	res := &errResponse{
		Err: err.Error(),
	}

	w.WriteHeader(code)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func sendJSON(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
}
