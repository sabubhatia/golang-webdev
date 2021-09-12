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
		log.Fatalf("Expected at least two args. Got: %d\n", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func dogs(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "Dogs", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func myDogs(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "MyDogs", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	f := http.StripPrefix("/public", fs)

	http.Handle("/public/", f)
	http.HandleFunc("/dogs/", dogs)
	http.HandleFunc("/mydogs/", myDogs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
