package main

import (
	"html/template"
	"log"
	"os"

	"github.com/sabubhatia/golang-webdev/01_templates/utility/fileutil"
)

type person struct {
	Fname string
	Lname string
}

type domicile struct {
	City    string
	Country string
}

var tpl *template.Template
var path string

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("Expected at least 2 args. Got: %d\n", len(os.Args))
	}

	tpl = template.Must(template.ParseGlob(os.Args[1]))
	log.Println(tpl.DefinedTemplates())

	// set the path from the environment.
	if len(os.Args) == 2 {
		return
	}

	path = os.Args[2]
}

func main() {
	// To do:
	//		struct, slice of struct, struct with slice of struct
	//
	// Think about:
	// How do I stucture my data in a composite data structure
	// How do I pass it in.
	// How do I access it.
	//
	// Anonymous structs can also be used
	//
	executeStruct("struct.gohtml")
	executeXStruct("structX.gohtml")
	executeStructWithX("structWithX.gohtml")
}

func executeStruct(tName string) {
	me := person{
		Fname: "Sabu",
		Lname: "Bhatia",
	}

	f := fileutil.OutF(path, tName)
	log.Printf("Will write too: %p, %v, %s", f, *f, f.Name())
	defer fileutil.CloseF(f)
	err := tpl.ExecuteTemplate(f, tName, me)
	if err != nil {
		log.Fatal("Unable to execute template struct.gohtml ", err)
	}
}

func executeXStruct(tName string) {
	px := personX()

	f := fileutil.OutF(path, tName)
	log.Printf("Will write too: %p, %v, %s", f, *f, f.Name())
	defer fileutil.CloseF(f)
	err := tpl.ExecuteTemplate(f, tName, px)
	if err != nil {
		log.Fatal("Unable to execute template structX.gohtml ", err)
	}
}

func executeStructWithX(tName string) {

	s := struct {
		People    []person
		Domiciles []domicile
	}{
		*personX(),
		*domicileX(),
	}
	if len(s.People) != len(s.Domiciles) {
		log.Fatalf("Expected #people %d to equial %d #domiciles ", len(s.People), len(s.Domiciles))
	}

	f := fileutil.OutF(path, tName)
	log.Printf("Will write too: %p, %v, %s", f, *f, f.Name())
	defer fileutil.CloseF(f)
	err := tpl.ExecuteTemplate(f, tName, s)
	if err != nil {
		log.Fatal("Unable to execute template structWithX.gohtml ", err)
	}
}

func personX() *[]person {
	px := []person{
		{
			Fname: "Sabu",
			Lname: "Bhatia",
		},
		{
			Fname: "Phengvanh",
			Lname: "Khounnhoth",
		},
		{
			Fname: "Kalwant",
			Lname: "Bhatia",
		},
	}

	return &px
}

func domicileX() *[]domicile {
	dx := []domicile{
		{
			City:    "Singapore",
			Country: "Singapore",
		},
		{
			City:    "Vientiane",
			Country: "Lao PDR",
		},
		{
			City:    "New Delhi",
			Country: "India",
		},
	}

	return &dx
}
