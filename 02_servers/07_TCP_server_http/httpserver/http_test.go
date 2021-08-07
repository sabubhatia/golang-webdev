package httpserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"testing"
	"text/template"
)

const (
	get = "GET"
	put = "PUT"
)

type request struct {
	action  string
	route   string
	version string
}

func Test_Http(t *testing.T) {
	startServer()
}

// Start server
// get conn
// parse request line.
// for get respond with a page.

func startServer() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)

	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	// Get the request-line
	// Get the header
	req, err := httpRequest(conn)
	if err != nil {
		log.Panic(err)
		return
	}
	log.Println("**Request line is: ", req)
	// Send a response if they issue a get()
	if req.action != get {
		log.Println("Exiting handle()")
		return
	}
	err = httpResponse(conn)
	if err != nil {
		log.Panic(err)
		return
	}
}

func httpRequest(conn net.Conn) (request, error) {

	scn := bufio.NewScanner(conn)
	// First line is is the request line:
	// action uri version
	i := 0
	var req request
	for scn.Scan() {
		txt := scn.Text()
		if txt == "" { //crlf received. End of header
			break
		}
		log.Println(txt)
		if i == 0 {
			if flds := strings.Fields(txt); len(flds) >= 3 {
				switch strings.ToUpper(flds[0]) {
				case "GET":
					req.action = get
				default:
					return req, fmt.Errorf("Action: %s, not recognised. ", flds[0])
				}
				req.route = flds[1]
				req.version = flds[2]
			} else {
				return req, fmt.Errorf("Bad request line: %s", txt)
			}
		}
		i++
	}

	log.Println("Exiting httpRequest()..")
	return req, nil
}

func httpResponse(conn net.Conn) error {
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

	tpl := template.Must(template.New("Respone").Parse(body))

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")

	fmt.Fprint(conn, body)
	// This is the response body..
	err := tpl.Execute(conn, tpl)
	if err != nil {
		return err
	}

	return nil
}
