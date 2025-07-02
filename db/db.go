package db

import (
	"database/sql" //std package
	"fmt"
	_ "modernc.org/sqlite" //sql db driver
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "events.db")
	if err != nil {
		panic("Error opening database:" + err.Error())
	}
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)
	createTable()
}

func createTable() {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	date_time DATETIME NOT NULL,
	user_id INTEGER NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`

	result, err := DB.Exec(createTableQuery)
	if err != nil {
		panic("Error creating table:" + err.Error())
	}
	fmt.Printf("%v", result)
}
