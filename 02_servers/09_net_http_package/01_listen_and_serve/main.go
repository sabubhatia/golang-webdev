package main

import (
	"fmt"
	"net/http"
)


type hotdog int


func (hotdog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Put anything here that you want...\n")
}
func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}