package main

import (
	"log"
	"math"
	"os"
	"text/template"

	"github.com/sabubhatia/golang-webdev/01_templates/utility/fileutil"
)

var fm = template.FuncMap{
	"fdbl":  dbl,
	"fsqr":  sqr,
	"fsqrt": sqrt,
}

var tpl *template.Template
var path = "./tmp/"

func init() {
	if len(os.Args) < 2 {
		log.Panicf("Expected at least 2 args but got: %d\n", len(os.Args))
	}
	tpl = template.Must(template.New("main").Funcs(fm).ParseGlob(os.Args[1]))
	log.Println(tpl.DefinedTemplates())
}

func dbl(v float64) float64 {
	return 2.0 * v
}

func sqr(v float64) float64 {
	return v * v
}

func sqrt(v float64) float64 {
	return math.Sqrt(v)
}

func main() {
	for _, t := range tpl.Templates() {
		f := fileutil.OutF(path, t.Name())
		defer fileutil.CloseF(f)
		err := tpl.ExecuteTemplate(f, t.Name(), 3.0)
		if err != nil {
			log.Panic("Unable to execute template: ", err)
		}
	}
}
