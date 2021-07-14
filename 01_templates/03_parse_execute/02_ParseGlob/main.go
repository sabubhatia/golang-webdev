package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	if len(os.Args) < 2 {
		log.Fatal("Expected at least 2 cli arguments Got: ", len(os.Args))
	}

	pattern := os.Args[1]
	tpl = template.Must(template.ParseGlob(pattern))
	log.Println(tpl.DefinedTemplates())
}

func outFile(fx []string) func() *os.File {

	// If no files are in fx return os.Stdout. Else keep returning teh next file in the list till
	// the end OF LIST IS HIT.

	useStdout := len(fx) <= 0
	next := 0
	ofx := fx
	return func() *os.File {
		if useStdout {
			return os.Stdout
		}
		var err error
		f, err := os.OpenFile(ofx[next], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal("Error during file open: ", err)
		}
		next += 1
		return f
	}
}

func main() {

	var ofx []string
	if len(os.Args) >= 3 {
		ofx = strings.Split(os.Args[2], ",")
		if len(ofx) < len(tpl.Templates()) {
			log.Fatalf("Expected: %d files for output, Got %d\n ", len(tpl.Templates()), len(ofx))
		}
	}

	getFile := outFile(ofx)
	for _, t := range tpl.Templates() {
		f := getFile()
		defer func() {
			if f == os.Stdout {
				return
			}
			log.Printf("Closing: %p, %v, %v\n", f, *f, f.Name())
			f.Close()
		}()
		err := t.ExecuteTemplate(f, t.Name(), nil)
		if err != nil {
			log.Fatal("Error executing template: ", t.Name(), err)
		}
	}
	log.Println("Done..")
}
