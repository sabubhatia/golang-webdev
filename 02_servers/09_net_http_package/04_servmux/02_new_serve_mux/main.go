package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1> Doggie Doggie !!</h1>")
}

type hotcat int

func (hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Kitty Kitty !!</h1>")
}
func main() {
	mux := http.NewServeMux()
	var d hotdog
	mux.Handle("/dog/", d)

	var c hotcat
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
