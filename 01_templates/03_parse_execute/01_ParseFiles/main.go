package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Expected 2 arguments got ", len(os.Args))
	}
	fx := strings.Split(os.Args[1], ",")
	tpl, err := template.ParseFiles(fx...)
	if err != nil {
		log.Fatal("Error during parse files: ", err)
	}

	var ofx []string
	if len(os.Args) > 2 {
		ofx = strings.Split(os.Args[2], ",")
	}

	if len(ofx) < len(fx) {
		log.Fatal("Expetced writers to equal readers got: ", len(ofx), len(fx))
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Unable to execute template: ", err)
	}

	for i, t := range tpl.Templates() {
		f, err := os.OpenFile(ofx[i], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal("Unable to open file for output: ", ofx[i], err)
		}
		defer f.Close()
		err = tpl.ExecuteTemplate(f, t.Name(), nil)
		if err != nil {
			log.Fatal(tpl.DefinedTemplates(), "\nUnable to execute template ", t.Name(), err)
		}
		// f.Close()
	}

	log.Println(tpl.DefinedTemplates())
}
