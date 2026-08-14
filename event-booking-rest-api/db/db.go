package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // _ means it will not be used directly, but it will be used under the hook by database/sql
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	const createEventsTable = `
	CREATE TABLE IF NOT EXISTS event (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  name TEXT NOT NULL,
	  description TEXT NOT NULL,
	  dateTime DATETIME NOT NULL,
	  user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table.")
	}
}
