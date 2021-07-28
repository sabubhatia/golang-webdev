package main

import (
	"log"
	"os"
	"text/template"

	"github.com/sabubhatia/golang-webdev/01_templates/13_hands_on/03/menu"
	"github.com/sabubhatia/golang-webdev/01_templates/utility/fileutil"
)

var tpl *template.Template

func init() {
	if len(os.Args) < 2 {
		log.Fatalln("Expected at least two arguments but got: ", len(os.Args))
	}

	tpl = template.Must(template.ParseGlob(os.Args[1]))
	log.Println(tpl.DefinedTemplates())
}

func main() {
	m := menu.Menus()

	f := fileutil.OutF("./tmp/", "Menus")
	defer fileutil.CloseF(f)
	err := tpl.ExecuteTemplate(f, "Menus", m)
	if err != nil {
		log.Panicf("Unable to execute template: %s", err.Error())
	}
}
