package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func Init(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbFile)
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
	_, err = db.Exec(schema)
	if err != nil {
		log.Println("Ошибка создания таблицы:", err)
		return nil, err
	}
	return db, nil
}
