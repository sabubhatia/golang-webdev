package main

import (
	"io"
	"log"
	"net/http"
)

func set(w http.ResponseWriter, req *http.Request) {
	log.Println("Method in set() is: ", req.Method)
	c := http.Cookie{
		Name:  "my-cookie",
		Value: "some-value",
	}
	http.SetCookie(w, &c)
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "Written cookie: <br>")
	io.WriteString(w, c.String())
	io.WriteString(w, "<br>")
}

func read(w http.ResponseWriter, req *http.Request) {
	log.Println("Method in read() is: ", req.Method)
	w.Header().Set("Content-Type", "text/html")
	c, err := req.Cookie("my-cookie")
	handleError(w, err)
	if err == nil {
		io.WriteString(w, c.String()+"<br>")
	} else {
		io.WriteString(w, "my-cookie="+"NOT FOUND"+"<br>")
	}
}

func handleError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
}
func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/Read", read)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
