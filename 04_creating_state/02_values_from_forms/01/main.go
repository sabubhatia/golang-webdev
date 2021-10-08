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

func foo(w http.ResponseWriter, req *http.Request) {

	stuff := req.FormValue("stuff")
	var method string
	if method = req.FormValue("method"); method == "" {
		method = "Get"
	}
	d := struct {
		Stuff  string
		Method string
	}{
		Stuff:  stuff,
		Method: method,
	}
	err := tpl.ExecuteTemplate(w, "Forms", d)
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
