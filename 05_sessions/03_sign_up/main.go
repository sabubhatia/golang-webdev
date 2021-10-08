package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type user struct {
	Username string
	First    string
	Last     string
	Pwd      string
}

var (
	dbUsers    = map[string]user{}   // un->user
	dbSessions = map[string]string{} // sid->un
	tpl        *template.Template
)

func init() {
	if len(os.Args) < 2 {
		log.Fatal("Expected at least 2 args. Got: ", len(os.Args))
	}

	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	err := tpl.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	if !alreadyLogedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	u := getUser(req)
	err := tpl.ExecuteTemplate(w, "bar.gohtml", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLogedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u user
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		if _, ok := dbUsers[un]; ok {
			// This username is already taken.
			http.Error(w, "Username is already taken. Use different username", http.StatusForbidden)
			return
		}
		f := req.FormValue("first")
		l := req.FormValue("last")
		p := req.FormValue("password")

		cookie, err := newsSessionCookie()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = un
		u = user{un, f, l, p}
		dbUsers[un] = u

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	err := tpl.ExecuteTemplate(w, "signup.gohtml", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func dump(w http.ResponseWriter, req *http.Request) {
	d := struct {
		DBSessions map[string]string
		DBUsers    map[string]user
	}{
		DBSessions: dbSessions,
		DBUsers:    dbUsers,
	}
	err := tpl.ExecuteTemplate(w, "dump.gohtml", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/dump", dump)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
