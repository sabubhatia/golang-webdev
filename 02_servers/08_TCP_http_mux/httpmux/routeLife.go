package httpmux

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"text/template"
)

type lifeHandler struct {
}

func NewLife() handleRoute {
	return &lifeHandler{}
}

func (*lifeHandler) Handle(conn net.Conn) error {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Life </title>
	</head>
	<body>
	<b> Life is: {{.}} </b>
	</body>
	</html>`

	w := bufio.NewWriter(conn)
	tpl := template.Must(template.New("Response").Parse(body))

	// This is the response body..
	err := tpl.ExecuteTemplate(w, "Response", "Like a box of chocolates..")
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", w.Size()-w.Available())
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	w.Flush()

	log.Println("Done..")
	return nil
}

func (*lifeHandler) String() string {
	return fmt.Sprintf("Life Handler)")
}
