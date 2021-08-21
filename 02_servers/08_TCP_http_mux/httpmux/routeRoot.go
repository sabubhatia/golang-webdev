package httpmux

import (
	"fmt"
	"html/template"
	"io"
)

type rootHandler struct {
}

func NewRoot() handleRoute {
	return &rootHandler{}
}

func (*rootHandler) String() string {
	return fmt.Sprintf("(Root Handler)")
}

func (*rootHandler) Body(w io.Writer) error {
	body := `<b> Hello, World ! </b>`

	tpl := template.Must(template.New("Response").Parse(body))

	// This is the response body..
	err := tpl.ExecuteTemplate(w, "Response", nil)
	if err != nil {
		return err
	}

	return nil
}

func (*rootHandler) Name() string {
	return fmt.Sprintf("Index")
}
