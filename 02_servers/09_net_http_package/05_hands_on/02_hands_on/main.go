package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

type data struct {
	Element string
	Form    string
}

func init() {
	if len(os.Args) < 2 {
		log.Fatal("Expected at least 2 arguments. Got: ", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" && req.Method != "" {
		log.Println("Handling Get index...")
		return
	}

	log.Println("Handling Get index...")
	// Handle the get method
	if err := tpl.ExecuteTemplate(w, "Index", data{"Index", ""}); err != nil {
		log.Panic(err)
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" && req.Method != "" {
		log.Println("Handling Get index...")
		return
	}

	log.Println("Handling Get dog...")
	// Handle the get method
	if err := tpl.ExecuteTemplate(w, "Index", data{"Dog", ""}); err != nil {
		log.Panic(err)
	}
}

func me(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" && req.Method != "" {
		log.Println("Bad method: ", req.Method)
		return
	}

	// Handle the get method
	log.Println("Handling Get me...")
	if err := tpl.ExecuteTemplate(w, "Index", data{"Me", "Me"}); err != nil {
		log.Panic(err)
	}
}

func inputName(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		log.Println("Bad method: ", req.Method)
		return
	}

	// Handle the post method
	log.Println("Handling Post inputName...")
	if err := req.ParseForm(); err != nil {
		log.Panic(err)
	}

	name := req.FormValue("name")
	if err := tpl.ExecuteTemplate(w, "Name", name); err != nil {
		log.Panic(err)
	}

}

func main() {

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/Dog/", http.HandlerFunc(dog))
	http.Handle("/Me", http.HandlerFunc(me))
	http.Handle("/Me/InputName", http.HandlerFunc(inputName))
	log.Println("Listening :8080....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
