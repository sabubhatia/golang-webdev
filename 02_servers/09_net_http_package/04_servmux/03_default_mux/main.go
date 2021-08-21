package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the big dog !")
}

type hotcat int

func (hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the little cat !")
}

func main() {

	var d hotdog
	var c hotcat

	// using the default mux.
	http.Handle("/dog/", d)
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil)
}
