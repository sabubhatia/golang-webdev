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
		log.Panicf("Expected at least 2 args but got: %d", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func main() {
	sx := []string{"MLK", "Gandhi", "Nanak", "Buddha", "Jesus", "Mohammed"}
	var fx []string
	if len(os.Args) >= 3 {
		fx = strings.Split(os.Args[2], ",")
	}
	getFile := filelist.FileList(fx)
	for _, t := range tpl.Templates() {
		f := getFile()
		defer func() {
			if f == os.Stdout {
				return
			}
			log.Printf("Closing: %p, %v, %v\n", f, *f, f.Name())
		}()
		err := t.ExecuteTemplate(f, t.Name(), sx)
		if err != nil {
			log.Panicf("Unable to execute template: %s %v\n", t.Name(), err)
		}
	}
}
