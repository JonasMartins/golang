package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var Db *sql.DB

func InitDB() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "dev", "_development", "hackernews")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Database successfully connected.")
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	m, err := migrate.New(
		"file://internal/pkg/db/migrations/postgresql",
		"postgres://dev:_development@localhost:5432/hackernews?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

}
