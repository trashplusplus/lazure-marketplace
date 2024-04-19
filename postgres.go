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

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASS")
	dbname := os.Getenv("DBNAME")
	dbport := os.Getenv("DB_PORT")

	log.Println("PORT: ", dbport)

	dbSource := " host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + dbport

	log.Println("DBSOURCE: ", dbSource)
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
