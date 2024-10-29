package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// SetupDatabase initializes SQLite database
func SetupDatabase(path string) {
	var err error
	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	createTables()
}

func createTables() {
	query := `
        CREATE TABLE IF NOT EXISTS schedules (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            url TEXT NOT NULL,
            start_time TEXT NOT NULL,
            end_time TEXT NOT NULL,
            day_of_week INTEGER
        );
    `
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}
