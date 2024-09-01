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

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env file %s", err.Error())
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASSWORD"),
		),
	)

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to database")
	}

	DB = db
	log.Println("Successfully connected to database")
}
