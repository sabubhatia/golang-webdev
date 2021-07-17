package main

import (
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/sabubhatia/golang-webdev/01_templates/utility/filelist"
)

var tpl *template.Template

func init() {
	if len(os.Args) < 2 {
		log.Panicf("Expected at least 2 arguemenst, Got: %d", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
	log.Println(tpl.DefinedTemplates())

}

func main() {

	var fx []string
	if len(os.Args) > 2 {
		fx = strings.Split(os.Args[2], ",")
	}

	getFile := filelist.FileList(fx)
	for _, t := range tpl.Templates() {
		f := getFile()
		defer func() {
			if f != os.Stdout {
				log.Printf("Closing: %p, %v, %v\n", f, *f, f.Name())
				f.Close()
			}

		}()
		err := t.ExecuteTemplate(f, t.Name(), "Life is like a box of chocolates...")
		if err != nil {
			log.Fatal("Error while executing template", err)
		}
	}

}