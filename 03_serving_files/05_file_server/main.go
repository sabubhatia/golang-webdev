package main

import (
	"io"
	"log"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	log.Println("Serving toby...")
	w.Header().Set("Content-type", "text/html;charset=utf-8")
	io.WriteString(w, "<h1>This is Toby...</h1><br>")
	io.WriteString(w, `<img src="Toby.jpg">`)
}

func root(w http.ResponseWriter, req *http.Request) {
	log.Println("Serving from root..")
	http.FileServer(http.Dir(".")).ServeHTTP(w, req)
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog", dog)
	log.Println("Listening :8080...")
	http.ListenAndServe(":8080", nil)
}
