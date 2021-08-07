package httpmux

import (
	"fmt"
	"io"
	"log"
	"text/template"
)

type lifeHandler struct {
}

func NewLife() handleRoute {
	return &lifeHandler{}
}

func (*lifeHandler) String() string {
	return fmt.Sprintf("Life Handler)")
}

func (*lifeHandler) Body(w io.Writer) error {
	body := `<b> Life is: {{.}} </b>` 
	
	tpl := template.Must(template.New("Response").Parse(body))

	// This is the response body..
	err := tpl.ExecuteTemplate(w, "Response", "Like a box of chocolates..")
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}


func (*lifeHandler) Name() string {
	return fmt.Sprintf("Life")
}
