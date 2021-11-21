package main

import (
	"fmt"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Fprintln(w, ctx)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
