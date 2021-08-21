package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"os"
	"strings"
	//"time"
)

var tpl *template.Template

type requestLine struct {
	method  string
	uri     string
	version string
}

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("Expected at least two args. Got: %d\n", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()
	log.Println("Listening tcp :8080...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	defer conn.Close()
	// conn.SetDeadline(time.Now().Add(20 * time.Second))

	var rl *requestLine
	first := true
	scn := bufio.NewScanner(conn)
	for scn.Scan() {
		txt := scn.Text()
		if txt == "" {
			log.Println("Empty string received. Breaking out...")
			break
		}
		log.Println("I got: " + txt)
		if first {
			first = false
			r, err := rLine(txt)
			if err != nil {
				log.Println(err)
				continue
			}
			rl = r
		}
	}
	if err := scn.Err(); err != nil {
		log.Println("Error:", err)
	}

	router(conn, *rl)
}

func rLine(txt string) (*requestLine, error) {
	sx := strings.Split(txt, " ")
	if len(sx) != 3 {
		return nil, fmt.Errorf("Not a HTTP request line...")
	}
	return &requestLine{sx[0], sx[1], sx[2]}, nil
}

func router(conn net.Conn, rl requestLine) {
	if rl.method == "GET" && rl.uri == "/" {
		index(conn)
		return
	}

	if rl.method == "GET" && rl.uri == "/Apply" {
		apply(conn)
		return
	}

	if rl.method == "POST" && rl.uri == "/ApplyForm" {
		applyForm(conn)
		return
	}
	notFound(conn)
}

func notFound(conn net.Conn) {
	io.WriteString(conn, "HTTP/1.1 404 Not Found Content\r\n")
	io.WriteString(conn, "Content-Type: text/html, charset=UTF-8\r\n")
	io.WriteString(conn, "\r\n")
}

func apply(conn net.Conn) {

	var sb strings.Builder
	if err := tpl.ExecuteTemplate(&sb, "Apply", nil); err != nil {
		log.Panic(err)
	}
	io.WriteString(conn, "HTTP/1.1 200\r\n")
	io.WriteString(conn, "Content-Type: text/html, charset=UTF-8\r\n")
	io.WriteString(conn, fmt.Sprintf("Content-length: %d\r\n", sb.Len()))
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, sb.String())
}

func applyForm(conn net.Conn) {
	r := "<p> Processing your received application form.. </p><br>"
	io.WriteString(conn, "HTTP/1.1 200\r\n")
	io.WriteString(conn, "Content-Type: text/html, charset=UTF-8\r\n")
	io.WriteString(conn, fmt.Sprintf("Content-length: %d\r\n", len(r)))
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, r)
}

func index(conn net.Conn) {
	var sb strings.Builder
	if err := tpl.ExecuteTemplate(&sb, "Index", nil); err != nil {
		log.Panic(err)
	}
	io.WriteString(conn, "HTTP/1.1 200\r\n")
	io.WriteString(conn, "Content-Type: text/html, charset=UTF-8\r\n")
	io.WriteString(conn, fmt.Sprintf("Content-length: %d\r\n", sb.Len()))
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, sb.String())
}

func main() {
	server()
}
