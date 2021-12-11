package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var (
	tpl *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
	log.Println(tpl.DefinedTemplates())
}

func index(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "index.gohtml", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is a message from foo()...")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
