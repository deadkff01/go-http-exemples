package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// var to store database
var DB *sql.DB

// make database connection
func Connect() {
	connStr := "user=postgres dbname=gogo password=123456 host=localhost sslmode=disable"
	con, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	DB = con
}

// close database connection
func Close() {
	DB.Close()
}
