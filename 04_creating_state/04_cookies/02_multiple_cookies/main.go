package main

import (
	//	"fmt"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func set(w http.ResponseWriter, req *http.Request) {
	c := http.Cookie{
		Name:  "My-Cookie",
		Value: "Some value",
		Path:  "/",
	}

	w.Header().Set("Content-Type", "text/html")
	http.SetCookie(w, &c)
	io.WriteString(w, "Set the cookie: "+c.String()+"<br>")
}

func read(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "Cookies received: <br>")
	for i, c := range req.Cookies() {
		io.WriteString(w, "Cookie#"+strconv.Itoa(i)+": "+c.String()+"<br>")
	}
}

func abundance(w http.ResponseWriter, req *http.Request) {
	m := map[string]string{
		"Dad":    "Kalwant",
		"Sister": "Rachna",
		"Wife":   "Pheng",
		"Me":     "Sabu",
	}

	log.Println("Received cookies: ", req.Cookies())
	w.Header().Set("Content-Type", "text/html>")
	for k, v := range m {
		c := http.Cookie{
			Name:  k,
			Value: v,
			Path:  "/",
		}
		http.SetCookie(w, &c)
	}
	io.WriteString(w, "Setting an abundance of cookies: <br>")
	io.WriteString(w, fmt.Sprint(w.Header())+"<br>")

	log.Println(w.Header())
}

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read/", read)
	http.HandleFunc("/abundance/", abundance)
	http.Handle("/favcicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
