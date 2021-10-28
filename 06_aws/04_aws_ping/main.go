package main

import (
	"io"
	"log"
	"net"
	"net/http"
)


func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "Hi! From your AWS server !")
}

func ping(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "Ping received !")
}

func instance(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "Instance on AWS is:  <br>")
	ax, err := net.InterfaceAddrs()
	if err != nil {
		io.WriteString(w, err.Error() + "<br>")
	}
	for _, a := range ax {
		io.WriteString(w, a.String() + "<br>")
	} 
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/instance", instance)
	http.HandleFunc("/ping", ping)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":80", nil))
}