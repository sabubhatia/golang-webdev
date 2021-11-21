package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
	Last  string
	Items []string
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset="UTF-8">
					<title> Foo </title>
				</head>
				<body>
					<p> This is the foo for JSON </p>
					<a href="/marshal"> Marshal </a> <br>
					<a href="/encode"> Encode </a> <br>
				</body>
			</html>
	`
	w.Write([]byte(s))
}

func marshal(w http.ResponseWriter, req *http.Request) {
	p := person{
		First: "James",
		Last:  "Bond",
		Items: []string{"suit", "car", "wry sense of humor"},
	}

	bs, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(bs)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(bs))
}

func encode(w http.ResponseWriter, req *http.Request) {
	p := person{
		First: "James",
		Last:  "Bond",
		Items: []string{"suit", "car", "wry sense of humor"},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil))
}
