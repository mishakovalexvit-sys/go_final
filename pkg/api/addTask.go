package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"main.go/pkg/db"
)

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
func checkDate(task *db.Task) error {
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format("20060102")
	}
	_, err := time.Parse("20060102", task.Date)
	if err != nil {
		return err
	}
	var next string
	if len(task.Repeat) > 0 {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}
	if now.Format("20060102") > task.Date {
		if len(task.Repeat) == 0 {
			task.Date = now.Format("20060102")
		} else {
			task.Date = next
		}
	}
	return nil
}
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "Ошибка десериализации JSON:" + err.Error()})
		return
	}
	if task.Title == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "Не указан заголовок задачи"})
		return
	}
	err = checkDate(&task)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	id, err := db.AddTask(&task)
	if err != nil {
		log.Println(err)
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "Ошибка сохранения в базу данных"})
		return
	}
	writeJson(w, http.StatusOK, map[string]any{"id": id})
}
