package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("Expected at least 2 args. Got: %d\n", len(os.Args))
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, "<p> Dog... </p>")
	io.WriteString(w, `<img src=/dog.jpg width=500 height=500>`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	log.Println("Serving Dog...")
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	io.Copy(w, f)

}

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/dog.jpg", dogPic)
	log.Println("Listening :8080")
	http.ListenAndServe(":8080", nil)
}
