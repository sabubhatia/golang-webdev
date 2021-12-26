package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

const (
	dsn = "postgres://bond:password@localhost:5432/bookstore?sslmode=disable"
	dn  = "pgx"
)

type book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	db, err := sql.Open(dn, dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()
	r, err := db.Query("select * from books")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	bks := make([]book, 0)
	for r.Next() {
		b := book{}
		if err := r.Scan(&b.isbn, &b.title, &b.author, &b.price); err != nil {
			panic(err)
		}
		bks = append(bks, b)
	}

	if err := r.Err(); err != nil {
		panic(err)
	}

	for _, b := range bks {
		fmt.Println(b)
	}
}
