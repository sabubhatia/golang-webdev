package main

import (
	"html/template"
	"io"
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
	log.Println("The method in foo is: ", req.Method)
	log.Println("Your name is : ", req.FormValue("first"))
}

func bar(w http.ResponseWriter, req *http.Request) {
	log.Println("The method in bar is: ", req.Method)
	f := req.FormValue("first")
	log.Println("Your name is: ", f)
	w.Header().Set("Location", "/?first="+f+" "+"Bhatia")
	w.WriteHeader(http.StatusSeeOther)
	io.WriteString(w, "Redirecting...")

}

func barred(w http.ResponseWriter, req *http.Request) {
	log.Println("The method in barred is: ", req.Method)

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
