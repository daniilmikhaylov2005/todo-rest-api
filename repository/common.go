package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func getConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://methil:301583@localhost:5432/todosdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
