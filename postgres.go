package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Loading env file error: %v", err)
	}

	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")
	dbname := os.Getenv("DBNAME")
	dbport := os.Getenv("DBPORT")

	dbSource := " host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + dbport

	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to the database!")
	return db, err
}
