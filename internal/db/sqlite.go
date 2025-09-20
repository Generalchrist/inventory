package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var DB *sqlx.DB

// Connect opens a SQLite database connection
func Connect() {
	var err error
	DB, err = sqlx.Open("sqlite3", "./inventory.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Test connection
	if err = DB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	log.Println("Connected to SQLite successfully ✅")
}
