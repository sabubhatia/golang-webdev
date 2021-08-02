package memcache

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"testing"
)

type cache map[string]string

var c cache = cache{}

func TestMemcache(t *testing.T) {
	startServer()
}

func startServer() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintln(conn, "Welcome to memory cache:")
	fmt.Fprintln(conn, usage())

	scn := bufio.NewScanner(conn)
	for scn.Scan() {
		txt := scn.Text()
		fs := strings.Fields(txt)
		if len(fs) < 1 {
			continue
		}

		switch cmd := strings.ToUpper(fs[0]); cmd {
		case "PUT":
			if len(fs) < 3 {
				fmt.Fprintln(conn, "Incorrect usage: ", txt)
				fmt.Fprintln(conn, usage())
				continue
			}
			c[fs[1]] = fs[2]
		case "GET":
			if len(fs) < 2 {
				fmt.Fprintln(conn, "Incorrect usage", txt)
				fmt.Fprintln(conn, usage())
				continue
			}
			if v, ok := c[fs[1]]; ok {
				fmt.Fprintln(conn, fs[1], ":", v)
			} else {
				fmt.Fprintln(conn, fs[1], "Does not exist")
			}
		case "COUNT":
			if len(c) < 1 {
				fmt.Fprintln(conn, "Cache is empty !")
				continue
			}
			fmt.Fprintf(conn, "Entries in cache = %d\n", len(c))
		case "DEL":
			if len(c) < 1 {
				continue
			}
			if len(fs) < 2 {
				fmt.Fprintln(conn, "Incorrect usage", txt)
				fmt.Fprintln(conn, usage())
				continue
			}
			delete(c, fs[1])
		case "LIST":
			if len(c) < 1 {
				fmt.Fprintln(conn, "Cache is empty !")
				continue
			}
			for k, v := range c {
				fmt.Fprintln(conn, k, ":", v)
			}
		default:
			fmt.Fprintln(conn, fs[0], " Not recognosed as a valid command")
			fmt.Fprintln(conn, usage())

		}
	}
}

func usage() string {
	return `Usage: [Put "key" "value] | [Get "key"] | [Del "key"] | [List] | [Count]`
}
