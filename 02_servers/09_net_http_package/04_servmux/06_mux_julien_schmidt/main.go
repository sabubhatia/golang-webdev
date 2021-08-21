package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	if len(os.Args) < 2 {
		log.Fatal("Expected at least 2 arguments. Got: ", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func index(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	s := struct {
		Element string
		Form    string
	}{
		"Index",
		"",
	}

	if err := tpl.ExecuteTemplate(w, "Index", s); err != nil {
		log.Panic(err)
	}
}

func readBlog(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	s := struct {
		Element string
		Form    string
	}{
		"Reading Blog",
		"",
	}

	if err := tpl.ExecuteTemplate(w, "Index", s); err != nil {
		log.Panic(err)
	}
}

func writeBlog(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	s := struct {
		Element string
		Form    string
	}{
		"Writng Blog",
		"",
	}

	if err := tpl.ExecuteTemplate(w, "Index", s); err != nil {
		log.Panic(err)
	}
}

func blogReader(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	s := struct {
		Element string
		Form    string
	}{
		"Blog Reader",
		"BlogReader",
	}

	if err := tpl.ExecuteTemplate(w, "Index", s); err != nil {
		log.Panic(err)
	}
}

func blogWriter(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	s := struct {
		Element string
		Form    string
	}{
		"Blog Writer",
		"BlogWriter",
	}

	if err := tpl.ExecuteTemplate(w, "Index", s); err != nil {
		log.Panic(err)
	}
}
func main() {

	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/Index", index)
	mux.GET("/BlogReader", blogReader)
	mux.GET("/BlogWriter", blogWriter)
	mux.GET("/ReadBlog", readBlog)
	mux.POST("/WriteBlog", writeBlog)
	log.Println("Listening on 8080..")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
