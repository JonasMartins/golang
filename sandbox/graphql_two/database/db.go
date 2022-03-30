package database

import (
	"database/sql"
	"graphql_two/utils"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	dns := os.Getenv("DNS")

	db, err := sql.Open("postgres", dns)

	utils.CheckError(err)

	defer db.Close()

	return db

}
