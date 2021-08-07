package httpmux

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"text/template"
)

type route struct {
	method string
	uri string
}

type registry map[route]handleRoute

var reg registry

func execute(conn net.Conn, method string, uri string) error {
	if uri == "" {
		return fmt.Errorf("Uri cannot be empty..")
	}
	loadRegistry()
	h := reg[route{method,uri}]
	if h == nil {
		return fmt.Errorf("Method: [%s], Route: [%s] not recognised ", method, uri)
	}
	return page(conn, h)
}

func loadRegistry() {
	if reg != nil {
		return
	}
	reg = make(registry)
	reg[route{get, "/"}] = NewRoot()
	reg[route{get, "/BALANCE",}] = NewBalance()
	reg[route{get, "/LIFE",}] = NewLife()
	reg[route{get, "/APPLY"}] = NewApply()
	reg[route{post, "/APPLY"}] = NewApplyProcess()

	log.Println("Elements in teh registry:\n", reg)
}

func page(conn net.Conn, h handleRoute) error {
	body := `
	<!DOCTYPE html>
	<html lang ="en">
	<head>
		<meta charset="utf-8">
		<title> Mux </title>
	</head>
	<body>
		<strong> {{.Head}} </strong><br>
		<a href="/">Index</a><br>
		<a href="/BALANCE">Balance</a><br>
		<a href="/Life">Life2</a><br>
		<a href="/Apply">Apply<a><br>
		<br><br><br>
		{{.Body}}
	</body>
	`

	var s strings.Builder

	err := h.Body(&s)
	if err != nil {
		return err
	}

	tpl := template.Must(template.New("Response").Parse(body))

	w := bufio.NewWriter(conn)
	// This is the response body..
	v := struct { 
		Head string
		Body string
	} {
		h.Name(),
		s.String(),
	}
	err = tpl.ExecuteTemplate(w, "Response", v)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", w.Size()-w.Available())
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	w.Flush()

	return nil
}
