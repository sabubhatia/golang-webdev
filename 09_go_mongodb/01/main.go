package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	http.ListenAndServe(":8080", mux)
}
