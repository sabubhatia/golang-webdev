package main

import (
	"html/template"
	"log"
	"os"
	"strings"
	"github.com/sabubhatia/golang-webdev/01_templates/utility/filelist"
)

var tpl *template.Template

func init() {
	if len (os.Args) < 2 {
		log.Fatal("Expected at least 2 args got ", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
	log.Println(tpl.DefinedTemplates())
}

func main() {

	var ofx []string
	if len(os.Args) >= 3 {
		ofx = strings.Split(os.Args[2], ",")
		if len(ofx) < len(tpl.Templates()) {
			log.Fatalf("Expected at least %d, Got %d ", len(tpl.Templates()), len(ofx))
		}
	}

	for _, t := range tpl.Templates() {
		t.ExecuteTemplate(f, t.Name(), "Meaning of life is 42")
	}
}