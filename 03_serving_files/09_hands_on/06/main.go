package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

var tpl *template.Template

type data struct {
	Element string
	Form    string
}

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("Expected at least two args. Got: %d\n", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func param(route string, form bool) data {
	var d data

	d.Element = route
	if form {
		d.Form = route
	}

	return d
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "Index", param("Index", false))
	handleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "Index", param("About", false))
	handleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "Index", param("Contact", false))
	handleError(w, err)
}

func applyProcess(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "Index", param("ApplyProcess", false))
	handleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		err := tpl.ExecuteTemplate(w, "Index", param("Apply", true))
		handleError(w, err)
		return
	}

	applyProcess(w, req)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
