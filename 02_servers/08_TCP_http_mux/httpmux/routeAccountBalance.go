package httpmux

import (
	"fmt"
	"io"
	"log"
	"text/template"
)

type balanceHandler struct {
}

func NewBalance() handleRoute {
	return &balanceHandler{}
}

func (*balanceHandler) String() string {
	return fmt.Sprintf("(Account Balance Handler)")
}

func (*balanceHandler) Body(w io.Writer) error {
	body := `<b> Your balance dear is: {{.}} </b>`

	tpl := template.Must(template.New("Response").Parse(body))

	// This is the response body..
	err := tpl.ExecuteTemplate(w, "Response", 42)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (*balanceHandler) Name() string {
	return fmt.Sprintf("Balance")
}
