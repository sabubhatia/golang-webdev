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

	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	req.Body.Close()
	f := req.FormValue("first")
	l := req.FormValue("last")
	e := req.FormValue("enct")
	if e == "" {
		log.Println("Setting default Enctype...")
		e = "application/x-www-form-urlencoded"
	}
	s := req.FormValue("subscribed") == "on"
	log.Println("Body: ", string(bs))
	d := struct {
		Enctype    string
		Body       string
		First      string
		Last       string
		Subscribed bool
	}{
		e,
		string(bs),
		f,
		l,
		s,
	}

	log.Println("d:", d, len(d.Body))
	err := tpl.ExecuteTemplate(w, "Form", d)
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	log.Println("Handling error...")
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
