package httpmux

import (
	"fmt"
	"net"
)

type handleRoute interface {
	fmt.Stringer
	Handle(con net.Conn) error
}
