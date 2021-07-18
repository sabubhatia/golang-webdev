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
	if len(os.Args) < 2 {
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

	getFile := filelist.FileList(ofx)
	for _, t := range tpl.Templates() {
		f := getFile()
		defer func() {
			if f != os.Stdout {
				log.Printf("Closing: %p, %v, %v\n", f, *f, f.Name())
				f.Close()
			}
		}()
		err := t.ExecuteTemplate(f, tpl.Name(), "Meaning of life is 42")
		if err != nil {
			log.Fatal("Unable to execute template", err)
		}
	}
}
