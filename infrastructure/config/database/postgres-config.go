package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var host = os.Getenv("DB_HOST")
	var user = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PASSWORD")
	var name = os.Getenv("DB_NAME")
	var schema = os.Getenv("DB_SCHEMA")

	connStr := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s search_path=%s sslmode=disable", host, user, password, name, schema)
	fmt.Println("connection: ", connStr)

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
