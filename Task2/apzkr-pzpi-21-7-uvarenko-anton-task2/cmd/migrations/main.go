package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	arg := os.Args[1]
	godotenv.Load()
	db, _ := sql.Open("mysql", os.Getenv("CONNECTION_STRING"))
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations/schema",
		"mysql",
		driver,
	)

	if arg == "up" {
		fmt.Println("going up migrations")
		err := m.Up()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
	if arg == "down" {
		fmt.Println("going down migrations")
		err := m.Down()
		fmt.Printf("%v\n", err)
	}
}
