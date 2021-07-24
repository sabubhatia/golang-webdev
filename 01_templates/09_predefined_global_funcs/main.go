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
		log.Fatalf("Expected at least 2 args got: %d\n", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
	log.Println(tpl.DefinedTemplates())
}

func main() {
	s := struct {
		Sages     map[string]string
		Countries []string
		Dumb      []string
	}{
		sageAndQuotes(),
		countries(),
		dumbStuff(),
	}

	for _, t := range tpl.Templates() {
		f := fileutil.OutF(path, t.Name())
		defer fileutil.CloseF(f)
		err := tpl.ExecuteTemplate(f, t.Name(), &s)
		if err != nil {
			log.Panic("Unable to execute template: ", err)
		}
	}
}

func sageAndQuotes() map[string]string {
	m := map[string]string{
		"Nanak": "Ek Onkar",
		"Jesus": "Lord have mercy",
		"Ram":   "Om",
		"Shiva": "Jai Shiv Shankar",
	}

	return m
}

func countries() []string {
	s := []string{"India", "Indonesia", "Japan", "Loas", "Malaysia", "Thailand", "USA"}

	return s
}

func dumbStuff() []string {
	s := []string{"Hi", "Bye", "", "Fun", "Boring", "Run", "Play"}

	return s
}
