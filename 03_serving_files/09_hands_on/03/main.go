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
		log.Fatalf("Expected at least 2 args. Got: %d", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func dog(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "Dog", nil); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func mydog(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "MyDog", nil); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.Handle("/pics/", fs)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/mydog/", mydog)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
