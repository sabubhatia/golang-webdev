package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch(req.URL.Path) {
	case "/dog":
		io.WriteString(w, "<h1> Doggie Doggie Doggie </h1>")
	case "/cat":
		io.WriteString(w, "<h1> Kittie, Kittie, Kittie </h1>")
	}
	
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}