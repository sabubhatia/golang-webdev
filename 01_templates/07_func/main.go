package main

import (
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/sabubhatia/golang-webdev/01_templates/utility/fileutil"
)

var fm = template.FuncMap{
	"up":    strings.ToUpper,
	"three": three,
}

var path = "./tmp"

var tpl *template.Template

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("Expecting at least 2 args got: %d \n", len(os.Args))
	}

	tpl = template.Must(template.New("sages_func.gohtml").Funcs(fm).ParseGlob(os.Args[1]))

	if len(os.Args) <= 2 {
		return
	}

	path = os.Args[2]
}

func main() {

	for _, t := range tpl.Templates() {
		f := fileutil.OutF(path, t.Name())
		defer fileutil.CloseF(f)
		err := tpl.ExecuteTemplate(f, t.Name(), sages())
		if err != nil {
			log.Fatal("Unable to excute(): ", err)
		}
	}
}

func three(s string) string {
	if len(s) >= 3 {
		return s[:3]
	}

	return s
}

func sages() *[]string {
	s := []string{
		"Gandhi",
		"Martin Luthor King",
		"Mohammed",
		"Jesus",
		"Nanak",
		"Buddha",
	}

	return &s
}
