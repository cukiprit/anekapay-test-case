package db

import (
	"database/sql"
	"log"
)

var Database *sql.DB

func InitDB() {
	var err error
	Database, err := sql.Open("sqlite3", "./db/animals.db")

	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS animals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		class TEXT,
		legs TEXT
	);
	`
	if _, err := Database.Exec(sqlStmt); err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
}

func CloseDB() {
	Database.Close()
}
