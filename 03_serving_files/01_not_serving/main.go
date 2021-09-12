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

func dog(w http.ResponseWriter, req *http.Request) {
	if err := tpl.ExecuteTemplate(w, "Index", nil); err != nil {
		log.Panic(err)
	}
	/*
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, `
		<!-----Not serving from our server------>
		<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
		`)
	*/
}

func dogDone(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<strong> Seen Dog </strong><br>`)
	io.WriteString(w, `<a href="/"> Back </a><br>`)
}
func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/DogDone", dogDone)
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
