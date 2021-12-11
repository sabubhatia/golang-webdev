package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

var (
	tpl *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func foo(w http.ResponseWriter, req *http.Request) {
	// Just write a mesage back..
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "Hi There ! This is a message from foo() !")
}

func index(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "index.gohtml", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
