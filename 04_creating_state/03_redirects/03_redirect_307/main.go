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
		log.Fatal("Expected at least 2 Args. Got: ", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func foo(w http.ResponseWriter, req *http.Request) {
	log.Println("Method in foo(): ", req.Method)
	log.Println("Your name in foo() is: ", req.FormValue("first"))
	log.Println("Your full name in foo() is: ", req.FormValue("fullname"))
}

func bar(w http.ResponseWriter, req *http.Request) {
	log.Println("Method in bar(): ", req.Method)
	log.Println("Your name in bar() is: ", req.FormValue("first"))
	// Do some processing...
	http.Redirect(w, req, "/?fullname="+req.FormValue("first")+" Bhatia", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, req *http.Request) {
	log.Println("Method in barred(): ", req.Method)
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
