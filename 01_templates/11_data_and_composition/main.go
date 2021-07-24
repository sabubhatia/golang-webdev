package main

import (
	"html/template"
	"log"
	"os"

	"github.com/sabubhatia/golang-webdev/01_templates/11_data_and_composition/course"
	"github.com/sabubhatia/golang-webdev/01_templates/utility/fileutil"
)


var tpl *template.Template
var path = "./tmp/"

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("Expecting at least 2 arguments but got: %d\n", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func main() {
	c := course.GetCourses(2021)
	f := fileutil.OutF(path, "Schedule")
	defer fileutil.CloseF(f)
	err := tpl.ExecuteTemplate(f, "Schedule", c)
	if err != nil {
		log.Panic("Unable to execute template: ", err)
	}
}