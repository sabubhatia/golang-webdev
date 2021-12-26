package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sabubhatia/golang-webdev/11_postgres/03_web_app/controller"
	"github.com/sabubhatia/golang-webdev/11_postgres/03_web_app/model"
	_ "github.com/jackc/pgx/stdlib"
)

var url = "postgres://bond:password@localhost:5432/" + model.DBName + "?sslmode=disable"

func getSession() *sql.DB {
	log.Println("Opening DB: " + url)
	db, err := sql.Open("pgx", url)
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	mux := httprouter.New()
	db := getSession()
	defer db.Close()
	bc := controller.NewBookController(db)
	mux.GET("/books", bc.GetBooks)
	log.Println("ListenAndServer() on :80...")
	log.Fatal(http.ListenAndServe(":80", mux))
}
