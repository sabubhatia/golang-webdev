package main

import (
	"text/template"
	"log"
	"os"
)

type input struct {
	Title string
	Heading string
	Input string
} 

func main() {
	s := input {
		Title: "Escaping",
		Heading: "Text template does not escape",
		Input: `<script>alert("You have been pwned !");</script>`,
	}

	tpl := template.Must(template.ParseGlob(os.Args[1]))
	err := tpl.Execute(os.Stdout, s)
	if err != nil {
		log.Panic("Unable to execute template: ", err.Error())
	}
} 