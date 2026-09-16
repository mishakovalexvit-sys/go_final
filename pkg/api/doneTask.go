package api

import (
	"net/http"
	"time"

	"main.go/pkg/db"
)

func DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("id")
	if id == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "Не указан идентификатор"})
		return
	}
	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "Задача не найдена"})
		return
	}
	if task.Repeat == "" {
		db.DeleteTask(id)
	} else {
		next, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		db.UpdateDate(next, id)
	}
	writeJson(w, http.StatusOK, map[string]any{})
}
