package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("Expected at least 2 args. Got: %d\n", len(os.Args))
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	io.WriteString(w, "<img src=/Dog.jpg>")
}

func getDog(w http.ResponseWriter, req *http.Request) {
	log.Println("Serving file..", os.Args[1])
	http.ServeFile(w, req, os.Args[1])
}

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/Dog.jpg", getDog)
	log.Println("Listening on :8080..........")
	http.ListenAndServe(":8080", nil)
}
