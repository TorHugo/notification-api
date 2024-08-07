package database

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func Start() {
	connStr := "host=localhost port=5432 user=admin password=admin dbname=notification sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %s", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	fmt.Println("Successfully connected to the database")
}

func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			return
		}
	}
}
