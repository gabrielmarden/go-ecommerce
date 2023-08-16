package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DbConnector() *sql.DB {
	connection := "user=dev dbname=postgres password=1234 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	return db
}
