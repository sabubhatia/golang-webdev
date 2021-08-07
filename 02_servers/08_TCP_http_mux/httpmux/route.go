package httpmux

import (
	"fmt"
	"io"
)

type handleRoute interface {
	fmt.Stringer
	Body(io.Writer) error
	Name() string
}
