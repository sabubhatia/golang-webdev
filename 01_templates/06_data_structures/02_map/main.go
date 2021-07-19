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
		log.Fatalf("Expected at least two command line arguments. Got: %d\n", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
	log.Println(tpl.DefinedTemplates())
}

func main() {

	var fx []string
	if len(os.Args) > 2 {
		fx = strings.Split(os.Args[2], ",")
		if len(fx) < len(tpl.Templates()) {
			log.Fatalf("Expected at least %d files. Got %d\n", len(tpl.Templates()), len(fx))
		}
	}
	getFile := filelist.FileList(fx)
	frx := map[string]string{
		"UK":        "Rob",
		"Canada":    "Greg",
		"NZ":        "Warren",
		"Singapore": "Rachpal",
		"Myanmar":   "Aung",
		"USA":       "Sabu",
	}

	for _, t := range tpl.Templates() {
		f := getFile()
		defer func() {
			if f == os.Stdout {
				return
			}
			log.Printf("Closing: %s\n", f.Name())
		}()
		err := tpl.ExecuteTemplate(f, t.Name(), frx)
		if err != nil {
			log.Panic("Error while executing template: ", t.Name(), " ", err)
		}
	}
}
