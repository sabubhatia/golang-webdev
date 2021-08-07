package httpmux

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"testing"
)

func Test_Mux(t *testing.T) {
	startServer()
}

func startServer() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) error {
	defer conn.Close()
	// get the request line" Method, URI and protocol version.
	hdr, err := getHeader(conn)
	if err != nil {
		return err
	}
	log.Printf("****Request line is: %v", hdr)

	// Satisfy the request by routing to appropriate route which may send the body
	err = execute(conn, hdr.Uri())
	if err != nil {
		log.Println("Unable to execute route: ", err)
	}

	return err
}

func getHeader(conn net.Conn) (Header, error) {
	// HEADER starts with the request line followed by the header fields and a final cr lf to indicate end of header
	// Method sp URI sp Protocol version is the request line we are after.

	scn := bufio.NewScanner(conn)
	var hdr Header
	i := 0
	for scn.Scan() {
		txt := scn.Text()
		log.Println(txt)
		if txt == "" {
			if i > 0 {
				return hdr, nil
			}
			return nil, fmt.Errorf("No request-line specified")
		}
		if i == 0 {
			if flds := strings.Fields(strings.ToUpper(txt)); len(flds) >= 3 {
				var method string
				switch flds[0] {
				case get:
					method = get
				default:
					return nil, fmt.Errorf("Inavlid request line. %s. %s unknown method", txt, flds[0])
				}

				hdr = New(method, flds[1], flds[2])
			} else {
				return nil, fmt.Errorf("Invalid request line: %s", txt)
			}
		}
		i++
	}
	return hdr, nil
}
