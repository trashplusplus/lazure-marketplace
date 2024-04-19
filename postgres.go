// here i need to create join to the database and include this file to the main.go
// postgresql
// create function initDB() to connect to the database and read config from .env file
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
	port := os.Getenv("PORT")

  log.Println("PORTTTT: ", port)

	dbSource := " host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port

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
