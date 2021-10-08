package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

type user struct {
	Username string
	First    string
	Last     string
}

var dbUsers = map[string]user{}      // map username to user
var dbSessions = map[string]string{} // map os sid to userna

func init() {
	if len(os.Args) < 2 {
		log.Fatal("Expected at least two args. Got ", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func index(w http.ResponseWriter, req *http.Request) {
	// get the cookie and the sid
	// if no sid create a session token
	// given sid get username
	// if no username found and a post request create user and asisgn username to sid
	// if username found and a post, override the association of sid to suername with username in post.

	cookie, err := req.Cookie("sid")
	if err != nil {
		// create a session token
		uid, err := uuid.NewV4()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// create a new cookie
		cookie = &http.Cookie{
			Name:     "sid",
			Value:    uid.String(),
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}
	}
	// Note: If request had a cookie the response just gets the same cookie back.
	http.SetCookie(w, cookie)
	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u, ok = dbUsers[un]
		if !ok {
			http.Error(w, "user for "+un+" not found", http.StatusInternalServerError)
		}
	}

	// at this point we have either a user in u or this session id has never been associated with a username hence it is empty.

	// Handle the post
	if req.Method == http.MethodPost {
		// get the form values.
		un := req.FormValue("username")
		f := req.FormValue("first")
		l := req.FormValue("last")
		u = user{
			Username: un,
			First:    f,
			Last:     l,
		}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u
	}

	err = tpl.ExecuteTemplate(w, "Index", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("sid")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u, ok = dbUsers[un]
		if !ok {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
	}

	err = tpl.ExecuteTemplate(w, "Bar", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func dump(w http.ResponseWriter, req *http.Request) {
	d := struct {
		DBUsers    map[string]user
		DBSessions map[string]string
	}{
		dbUsers,
		dbSessions,
	}
	err := tpl.ExecuteTemplate(w, "Dump", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/dump", dump)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
