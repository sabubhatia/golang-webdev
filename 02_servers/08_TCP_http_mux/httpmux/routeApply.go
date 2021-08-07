package httpmux

import (
	"fmt"
	"io"
	"log"
	"text/template"
)

type applyHandler struct {
}

func NewApply() handleRoute {
	return &applyHandler{}
}

func (*applyHandler) String() string {
	return fmt.Sprintf("Apply Handler)")
}

func (*applyHandler) Body(w io.Writer) error {
	body := `
	<form method="post" action="/apply">
	<input type="submit" value="apply">
	` 
	
	tpl := template.Must(template.New("Response").Parse(body))

	// This is the response body..
	err := tpl.ExecuteTemplate(w, "Response", nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}


func (*applyHandler) Name() string {
	return fmt.Sprintf("Apply")
}
