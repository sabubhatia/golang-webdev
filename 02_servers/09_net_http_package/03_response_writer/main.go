package main

import (
	"fmt"
	"log"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling...")
	w.Header().Set("Sabu-key", "This is sabu & pheng love key")
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	fmt.Fprintf(w, "<p><strong>You can show your love right here amigo !</strong></p>")
}

func main() {
	var d hotdog

	http.ListenAndServe(":8080", d)
}
