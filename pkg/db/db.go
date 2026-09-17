package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init(dbFile string) (*sql.DB, error) {
	var err error
	DB, err = sql.Open("sqlite", dbFile)
	if err != nil {
		log.Println("Ошибка запуска драйвера:", err)
		return nil, err
	}
	schema := `
	CREATE TABLE IF NOT EXISTS scheduler (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	date CHAR(8) NOT NULL DEFAULT "",
    	comment TEXT NOT NULL DEFAULT "",
		title VARCHAR NOT NULL DEFAULT "",
		repeat VARCHAR NOT NULL DEFAULT ""
	);`
	_, err = DB.Exec(schema)
	if err != nil {
		log.Println("Ошибка создания таблицы:", err)
		return nil, err
	}
	return DB, nil
}
