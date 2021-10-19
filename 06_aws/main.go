package main

import (
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "Yes! I am running on aws !")
}

func main() {
	http.HandleFunc("/", index)
	log.Println("Starting http server...")
	log.Println(http.ListenAndServe(":80", nil))
}