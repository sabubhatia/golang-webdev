package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("Exepcted at least 2 arguments. Got %d", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func dogs(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "Dogs", nil); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func myDogs(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "MyDogs", nil); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	f := http.StripPrefix("/resources", fs)

	http.Handle("/", fs)
	http.Handle("/resources/", f)
	http.HandleFunc("/dogs/", dogs)
	http.HandleFunc("/mydogs/", myDogs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
