package httpmux

import "fmt"

const (
	get = "GET"
	put = "PUT"
)

type headerStruct struct {
	method  string
	uri     string
	version string
}

type Header interface {
	String() string
	Method() string
	Uri() string
	Version() string
}

func (h *headerStruct) String() string {
	return fmt.Sprintf("Method (%s), Uri (%s), Version (%s)", h.method, h.uri, h.version)
}

func (h *headerStruct) Method() string {
	return h.method
}

func (h *headerStruct) Uri() string {
	return h.uri
}

func (h *headerStruct) Version() string {
	return h.version
}

func New(method, uri, version string) Header {
	return &headerStruct{
		method:  method,
		uri:     uri,
		version: version,
	}
}
