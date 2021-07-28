package main

import (
	"log"
	"os"
	"text/template"

	"github.com/sabubhatia/golang-webdev/01_templates/utility/fileutil"
)

type course struct {
	Number string
	Name   string
	Units  string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcaYear string
	Fall    semester
	Spring  semester
	Summer  semester
}

var tpl *template.Template

func init() {
	if len(os.Args) < 2 {
		log.Fatal("Expected at least 2 args got: ", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func main() {
	years := []year{
		{
			AcaYear: "2020-2021",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					{"CSCI-40", "Introduction to Programming in Go", "4"},
					{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					{"CSCI-50", "Advanced Go", "5"},
					{"CSCI-190", "Advanced Web Programming with Go", "5"},
					{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
		{
			AcaYear: "2021-2022",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					{"CSCI-40", "Introduction to Programming in Go", "4"},
					{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					{"CSCI-50", "Advanced Go", "5"},
					{"CSCI-190", "Advanced Web Programming with Go", "5"},
					{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
	}

	f := fileutil.OutF("./tmp/", "Catalog")
	defer fileutil.CloseF(f)

	err := tpl.ExecuteTemplate(f, "Catalog", years)
	if err != nil {
		log.Fatalln(err)
	}
}
