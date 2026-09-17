package main

import (
	"log"
	"net/http"

	"main.go/pkg/api"
	db "main.go/pkg/db"
)

func main() {
	dbFile := "scheduler.db"
	database, err := db.Init(dbFile)
	if err != nil {
		log.Fatal("Ошибка инициализации БД:", err)
	}
	defer database.Close()
	api.Init()
	port := "7540"
	http.Handle("/", http.FileServer(http.Dir("web")))
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Ошибка открытия порта:", err)
	}

}
