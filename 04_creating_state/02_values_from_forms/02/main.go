package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	First      string
	Last       string
	Subscribed bool
}

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("Expected at least 2 arguments. Received %d", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func foo(w http.ResponseWriter, req *http.Request) {
	f := req.FormValue("First")
	l := req.FormValue("Last")
	s := req.FormValue("Subscribed") == "on"
	m := req.FormValue("Method")
	if m == "" {
		m = "Get"
	}

	d := struct {
		Method string
		person
	}{
		m,
		person{f, l, s},
	}
	err := tpl.ExecuteTemplate(w, "Index", d)
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
