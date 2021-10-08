package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func count(w http.ResponseWriter, req *http.Request) {
	// Get the cookie called "visits"

	c, err := req.Cookie("visits")
	cnt := 0
	if err == nil {
		cnt, err = strconv.Atoi(c.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	cnt++
	w.Header().Set("Content-Type", "text/html")
	c = &http.Cookie{
		Name:  "visits",
		Value: strconv.Itoa(cnt),
		Path:  "/",
	}
	http.SetCookie(w, c)
	io.WriteString(w, "You have visited this domain "+strconv.Itoa(cnt)+" time(s) <br>")
}

func main() {
	http.HandleFunc("/", count)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
