package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB

var host = os.Getenv("DB_HOST")
var port = os.Getenv("DB_PORT")
var user = os.Getenv("DB_USER")
var password = os.Getenv("DB_PASSWORD")
var name = os.Getenv("DB_NAME")
var schema = os.Getenv("DB_SCHEMA")

func Start() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable", host, port, user, password, name, schema)
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
