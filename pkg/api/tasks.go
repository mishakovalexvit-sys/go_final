package api

import (
	"net/http"

	"main.go/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50) // в параметре максимальное количество записей
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "ошибка функци tasks"})
		return
	}
	writeJson(w, http.StatusOK, TasksResp{
		Tasks: tasks,
	})
}
