package main

import (
	"io"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func bar(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "In bar() The cookie is: "+cookie.String()+"<br>")
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		// session ID does not exist. generate a UUID

		uid, err := uuid.NewV4()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println(uid.String())
		cookie = &http.Cookie{
			Name:     "session",
			Value:    uid.String(),
			Secure:   true,
			HttpOnly: true,
			MaxAge:   30,
		}
		http.SetCookie(w, cookie)
	}
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "The cookie is: "+cookie.String()+"<br>")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
