package main

import (
	"log"
	"net/http"
)

func notFound(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Not found !", http.StatusNotFound)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/resources/", notFound)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
