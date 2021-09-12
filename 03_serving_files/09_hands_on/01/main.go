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
		log.Fatalf("Expected at least two args. Got: %d\n", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func foo(w http.ResponseWriter, req *http.Request) {
	log.Println("Handling foo...")
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	io.WriteString(w, "<strong> Foo ran ! </strong> <br>")
}

func dog(w http.ResponseWriter, req *http.Request) {
	log.Println("Handling dog...")
	if err := tpl.ExecuteTemplate(w, "Dog", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func serveToby(w http.ResponseWriter, req *http.Request) {
	log.Println("Handling serveToby..")
	http.ServeFile(w, req, "Toby.jpg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/resources/Toby.jpg", serveToby)
	log.Println("Listening :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
