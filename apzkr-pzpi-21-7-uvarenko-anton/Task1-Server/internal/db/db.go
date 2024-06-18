package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	connStr := os.Getenv("CONNECTION_STRING")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}
	return db
}
