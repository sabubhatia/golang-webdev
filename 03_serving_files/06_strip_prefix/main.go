package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	io.WriteString(w, "<strong>Here is Mr. Toby from strip prefix..<strong><br>")
	io.WriteString(w, `<img  src=/resources/toby.jpg>`)
}

func main() {

	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}
