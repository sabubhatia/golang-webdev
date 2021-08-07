package httpmux

import (
	"fmt"
	"html/template"
	"net"
)

type rootHandler struct {
}

func NewRoot() handleRoute {
	return &rootHandler{}
}

func (*rootHandler) Handle(conn net.Conn) error {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello, World </title>
	</head>
	<body>
	<b> Hello, World ! </b>
	</body>
	</html>`

	tpl := template.Must(template.New("Response").Parse(body))

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")

	// This is the response body..
	err := tpl.ExecuteTemplate(conn, "Response", nil)
	if err != nil {
		return err
	}

	return nil
}

func (*rootHandler) String() string {
	return fmt.Sprintf("(Root Handler)")
}
