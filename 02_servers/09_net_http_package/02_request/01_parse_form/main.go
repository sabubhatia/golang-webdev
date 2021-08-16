package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {

	if len(os.Args) < 2 {
		log.Fatal("Expected at least 2 arguments but got: ", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}


type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	err = tpl.ExecuteTemplate(w, "Input", r.Form)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}