package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
)

type prata int

var tpl *template.Template

func init() {
	if len(os.Args) < 2 {
		log.Fatalln("Expected at least two arguments. Got: ", len(os.Args))
	}

	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func (prata) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Panic(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
		URL         *url.URL
	}{
		req.Method,
		req.PostForm,
		req.URL,
	}

	if err := tpl.ExecuteTemplate(w, "URL", data); err != nil {
		log.Panic(err)
	}
}

func main() {
	var p prata
	http.ListenAndServe(":8080", p)
}
