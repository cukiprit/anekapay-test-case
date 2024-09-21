package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func InitDB() {
	var err error
	Database, err = sql.Open("sqlite3", "../internal/db/animals.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
    CREATE TABLE IF NOT EXISTS animals (
        id INTEGER PRIMARY KEY,
        name TEXT,
        class TEXT,
        legs INTEGER
    );
    `
	if _, err := Database.Exec(sqlStmt); err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
}

func CloseDB() {
	Database.Close()
}
