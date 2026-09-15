package api

import (
	"net/http"

	"main.go/pkg/db"
)
func GetTaskHandler(w http.ResponseWriter, r *http.Request){
	idStr := r.FormValue("id")
	if idStr == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "Не указан идентификатор"})
		return
	}
	t, err := db.GetTask(idStr)
	if err != nil{
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	writeJson(w, http.StatusOK, t)
}