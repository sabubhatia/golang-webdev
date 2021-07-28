package main

import (
	"log"
	"os"
	"text/template"

	"github.com/sabubhatia/golang-webdev/01_templates/utility/fileutil"
)

var tpl *template.Template
var path = "./tmp/"

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("Expected at least 2 args but got: %d\n", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func main() {
	f := fileutil.OutF(path, "Index.doc")
	defer fileutil.CloseF(f)
	err := tpl.ExecuteTemplate(f, "Index", nil)
	if err != nil {
		log.Fatal("Unable to execute template: ", err)
	}

	err = tpl.ExecuteTemplate(f, "Odd", nil)
	if err != nil {
		log.Fatal("Unable to execute template: ", err)
	}

	err = tpl.ExecuteTemplate(f, "Even", nil)
	if err != nil {
		log.Fatal("Unable to execute template: ", err)
	}

	f = fileutil.OutF(path, "Genius")
	defer fileutil.CloseF(f)
	s := struct {
		Footer string
		Header string
		Body   []string
	}{
		Footer: "This is a footer",
		Header: "This is a header",
		Body: []string{
			"Laos",
			"Thailand",
			"India",
			"Phillipines",
		},
	}
	err = tpl.ExecuteTemplate(f, "Genius", s)
	if err != nil {
		log.Fatal("Unable to execute template: ", err)
	}
}
