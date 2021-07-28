package main

import (
	"html/template"
	"log"
	"os"

	"github.com/sabubhatia/golang-webdev/01_templates/utility/fileutil"
)

var tpl *template.Template
var path = "./tmp/"

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("Expected at least 2 arguments. Got: %d\n", len(os.Args))
	}

	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

const (
	Male int = iota
	Female
)

type person struct {
	name string
	age  uint8
	sex  int
}

func (p person) Name() string {
	return p.name
}

func (p person) Age() uint8 {
	return p.age
}

func (p person) Sex() int {
	return p.sex
}

func (p person) IsMale() bool {
	return p.sex == Male
}

func (p person) DblAge() uint8 {
	return p.age * 2
}

func (p *person) NameChange(name string) string {
	if len(name) < 1 {
		log.Panicf("Name must contain at least one character.")
	}

	p.name = name

	return name
}

func main() {

	f := fileutil.OutF(path, "Methods")
	defer fileutil.CloseF(f)
	err := tpl.ExecuteTemplate(f, "Methods", &person{"Sabu Bhatia", 52, Male})
	if err != nil {
		log.Panic("Unable to execute template: ", err)
	}
}
