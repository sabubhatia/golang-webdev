package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, rsp *http.Request) {
	io.WriteString(w, "This is the dog handle func")
}

func c(w http.ResponseWriter, rsp *http.Request) {
	io.WriteString(w, "This is the cat handle func")
}

func main() {
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat/", c)
	http.ListenAndServe(":8080", nil)
}