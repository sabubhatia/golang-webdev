package httpmux

import (
	"fmt"
	"net"
)

type registry map[string]handleRoute

var reg registry

func execute(conn net.Conn, uri string) error {
	if uri == "" {
		return fmt.Errorf("Uri cannot be empty..")
	}
	loadRegistry()
	h := reg[uri]
	if h == nil {
		return fmt.Errorf("Route: [%s] not recognised ", uri)
	}
	return h.Handle(conn)
}

func loadRegistry() {
	if reg != nil {
		return
	}
	reg = make(registry)
	reg["/"] = NewRoot()
	reg["/BALANCE"] = NewBalance()
	reg["/LIFE"] = NewLife()
}
