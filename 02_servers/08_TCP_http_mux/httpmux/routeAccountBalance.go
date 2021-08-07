package httpmux

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"text/template"
)

type balanceHandler struct {
}

func NewBalance() handleRoute {
	return &balanceHandler{}
}

func (*balanceHandler) Handle(conn net.Conn) error {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Balance </title>
	</head>
	<body>
	<b> Your balance dear is: {{.}} </b>
	</body>
	</html>`

	w := bufio.NewWriter(conn)
	tpl := template.Must(template.New("Response").Parse(body))

	// This is the response body..
	err := tpl.ExecuteTemplate(w, "Response", 42)
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

func (*balanceHandler) String() string {
	return fmt.Sprintf("(Account Balance Handler)")
}
