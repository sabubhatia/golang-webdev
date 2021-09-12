package main

import (
	"io"
	"log"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	io.WriteString(w, "Look at your terminal..")
}

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/cat", http.NotFoundHandler())
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
