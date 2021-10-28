package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // _ is the alias a blank identifer. I will never use any code from here. Imported for set up
)

const (
	dbDriver = "mysql"
	dataSourceName = "admin:password@tcp(database-1.cauehzhs4qsn.ap-southeast-1.rds.amazonaws.com:3306)/test01?charset=utf8"
)

func index(w http.ResponseWriter, req *http.Request) {
	log.Println("Handling index()...")
	w.Header().Set("Content-Type", "text/html")
	log.Println("Opening DB")
	db, err := sql.Open(dbDriver, dataSourceName)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	defer db.Close()
	log.Println("DB opened ")
	if err := db.Ping(); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 	
	}
	
	io.WriteString(w, "Hi I have accessed the db on AWS !")
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Listening on Port :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}