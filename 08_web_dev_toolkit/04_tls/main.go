package main

import (
	"io"
	"log"
	"net/http"
)

// TLS exhanges an asymmetric key.
// Client generated a symetrix key and trasnmits that to server by encrypting with the public key above.
// Client server then use symetric key encryption to communicate
// Client server
// https ports are 10 and 443. 443 is production

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "This is an HTTPS server..My first one ever !")
}
func main() {
	http.HandleFunc("/", foo)
	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil))
}
