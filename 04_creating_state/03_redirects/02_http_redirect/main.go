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
		log.Fatal("Expected at least 2 args. Got: ", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func foo(w http.ResponseWriter, req *http.Request) {
	log.Println("Method in foo() is: ", req.Method)
	log.Println("Your first name in foo() is: ", req.FormValue("first"))
	log.Println("Your name in foo() is: ", req.FormValue("name"))
}

func bar(w http.ResponseWriter, req *http.Request) {
	f := req.FormValue("first")
	log.Println("Method in bar() is: ", req.Method)
	log.Println("Your name in bar() is: ", f)
	// Process form here..
	http.Redirect(w, req, "/?name="+f+" Bhatia", http.StatusSeeOther)
}

func barred(w http.ResponseWriter, req *http.Request) {
	log.Println("Method is barred() is: ", req.Method)
	err := tpl.ExecuteTemplate(w, "Index", nil)
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
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
