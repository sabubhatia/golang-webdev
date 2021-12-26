package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	url := "postgres://bond:password@localhost:5432/bookstore?sslmode=disable"
	db, err := sql.Open("pgx", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You are connected to the database")
}
