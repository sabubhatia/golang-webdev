package main

import (
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	log.Println("Method in foo(): ", req.Method)
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<p> Moved here permanently !! </p>")
}

func barmoved(w http.ResponseWriter, req *http.Request) {
	log.Println("Method in barmoved(): ", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/barmoved", barmoved)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
