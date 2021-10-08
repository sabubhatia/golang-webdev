package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, req *http.Request) {
	log.Println("In index...")
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<h1> <a href="/set"> Set </a></h1><br>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	log.Println("In set()...")
	cookie, err := req.Cookie("session")
	cnt := 0
	if err == nil {
		// Cookie exists.
		cnt, err = strconv.Atoi(cookie.Value)
		if err != nil {
			// cookie seems corrupted.
			http.Redirect(w, req, "/delete", http.StatusSeeOther)
			return
		}
	} else {
		log.Println("Cookie not found: ", req.URL)
		cookie = &http.Cookie{
			Name:  "session",
			Value: strconv.Itoa(cnt),
			Path:  "/",
		}
	}
	cnt++
	cookie.Value = strconv.Itoa(cnt)
	w.Header().Set("Content-Type", "text/html")
	http.SetCookie(w, cookie)
	io.WriteString(w, "Set cookie: "+cookie.String()+"<br>")
	io.WriteString(w, `<h1> <a href="/read"> Read </a></h1><br>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	log.Println("In read()...")
	cookie, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		log.Println("func read() Cookie not found..redirecting..")
		// no cookie found redirect to go back to set.
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "You have so far visited /set "+cookie.Value+" time(s) <br>")
	io.WriteString(w, `<h1> <a href="/"> Index </a></h1><br>`)
	io.WriteString(w, `<h1> <a href="/delete"> delete </a></h1><br>`)
}

func delete(w http.ResponseWriter, req *http.Request) {
	log.Println("In delete()...")
	cookie, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		// no cookie found redirect to go back to set.
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	// expire this cookie.
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "text/html")
	http.Redirect(w, req, "/set", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/delete", delete)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
