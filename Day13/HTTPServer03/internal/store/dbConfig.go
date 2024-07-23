package store

import (
	"database/sql"
	"log"
)

// global variable database
var database *sql.DB

// Connect to the "northwind_db" database
func DBConnect() {
	//connection string
	connStr := "user=postgres password=admin dbname=batch30_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	//state status koneksi ke db
	database = db
	log.Println("Database Connected.")
}
