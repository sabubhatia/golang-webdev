package main

import (
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "Search query:"+v+"\n")
	io.WriteString(w, "All passed in query parameters: \n")
	for k, v := range req.Form {
		io.WriteString(w, k+": ")
		fv := true
		for _, s := range v {
			if !fv {
				io.WriteString(w, ", ")
			}
			io.WriteString(w, s)
			fv = false
		}
		io.WriteString(w, "\n")
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
