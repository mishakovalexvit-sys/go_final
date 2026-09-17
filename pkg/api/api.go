package api

import (
	"net/http"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		AddTaskHandler(w, r)
	case http.MethodGet:
		GetTaskHandler(w, r)
	case http.MethodPut:
		UpdateTaskHandler(w, r)
	case http.MethodDelete:
		DeleteTaskHandler(w, r)
	default:
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "Метод не поддерживается"})
	}
}

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", DoneTaskHandler)
}
