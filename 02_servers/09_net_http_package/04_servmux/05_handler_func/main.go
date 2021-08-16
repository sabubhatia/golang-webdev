package main

import (
	"io"
	"net/http"
)


func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is a dog !")
}

func  c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is a cat !")
}


func dFunc(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is a dog FUNC!")
}

func  cFunc(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is a cat FUNC!")
}
func main() {


	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)

	http.Handle("/dog/func", http.HandlerFunc(dFunc))
	http.Handle("/cat/func", http.HandlerFunc(cFunc))
	http.ListenAndServe(":8080", nil)

}