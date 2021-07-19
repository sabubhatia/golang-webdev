package main

import (
	"html/template"
	"log"
	"os"
	"strings"
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

	// set the path from teh environment.
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

	f := outF(tName)
	log.Printf("Will write too: %p, %v, %s", f, *f, f.Name())
	defer closeF(f)
	err := tpl.ExecuteTemplate(f, tName, me)
	if err != nil {
		log.Fatal("Unable to execute template struct.gohtml ", err)
	}
}

func executeXStruct(tName string) {
	px := personX()

	f := outF(tName)
	log.Printf("Will write too: %p, %v, %s", f, *f, f.Name())
	defer closeF(f)
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

	f := outF(tName)
	log.Printf("Will write too: %p, %v, %s", f, *f, f.Name())
	defer closeF(f)
	err := tpl.ExecuteTemplate(f, tName, s)
	if err != nil {
		log.Fatal("Unable to execute template structWithX.gohtml ", err)
	}
}

func outF(tn string) *os.File {
	if len(path) < 1 {
		return os.Stdout
	}

	sx := strings.Split(tn, ".")
	if len(sx) != 2 {
		log.Panicf("Expected %s to split into 2 but got: %d", tn, len(sx))
	}
	fn := strings.Join([]string{path, sx[0]}, "")
	fn = strings.Join([]string{fn, "html"}, ".")
	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Panic("Unable to open: ", fn, " ", err)
	}
	return f
}

func closeF(f *os.File) {
	if f == os.Stdout {
		return
	}

	log.Printf("Closing: %p, %v, %s", f, *f, f.Name())
	f.Close()
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
