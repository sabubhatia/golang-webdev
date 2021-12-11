package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Un    string
	First string
	Last  string
	Pwd   []byte
	Role  string
}

type session struct {
	Un           string
	LastActivity time.Time
}

var (
	dbUsers    = map[string]user{}    // User name -> user
	dbSessions = map[string]session{} // Session ID -> session
	tpl        *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
	log.Println(tpl.DefinedTemplates())
	addUsers()
}

func addUsers() {
	// Prime db users db. Again thsi is just for demo. In real life have a proper DB.
	p, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	u := user{
		Un:    "sabu@gmail.com",
		First: "Sabu",
		Last:  "Singh",
		Pwd:   p,
		Role:  "Admin",
	}

	dbUsers[u.Un] = u
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	d := struct {
		Un  string
		Sid string
		La  string
	}{
		"",
		"",
		"",
	}
	if cookie, err := req.Cookie(sidName); err == nil {
		if s, ok := dbSessions[cookie.Value]; ok {
			d.Un = u.Un
			d.Sid = cookie.Value
			d.La = fmtTime(s.LastActivity)
		} 
	}

	if err := tpl.ExecuteTemplate(w, "index.gohtml", d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func userExists(w http.ResponseWriter, req *http.Request) {
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	un := string(bs)
	_, ok := dbUsers[un]
	log.Println("[", un, "]", ok)
	fmt.Fprint(w, ok)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "Must be 007 to enter the bar. Your role is: "+u.Role, http.StatusForbidden)
		return
	}

	if err := tpl.ExecuteTemplate(w, "bar.gohtml", u); err != nil {
		return
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		return
	}

	if req.Method == http.MethodPost {
		// Get the username and pwd
		un := req.FormValue("username")

		// check if the user exists
		if _, ok := dbUsers[un]; !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// If exists bcrypt password and compare to existing password
		pwd := req.FormValue("pwd")
		if bcrypt.CompareHashAndPassword(dbUsers[un].Pwd, []byte(pwd)) != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// if match login and create a sesssion
		sid := createSession(un)
		cookie := &http.Cookie{
			Name:   sidName,
			Value:  sid,
			MaxAge: maxAge,
		}

		http.SetCookie(w, cookie)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		log.Println("Session ", sid, " Created.", "Active sessions: ", len(dbSessions))
		return
	}

	if err := tpl.ExecuteTemplate(w, "login.gohtml", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		return
	}

	cookie, err := req.Cookie(sidName)
	if err != nil {
		//This cannot happen. Since we are logged in a cookie must exist. So return error.
		http.Error(w, "Cookie not found. Unexpected "+err.Error(), http.StatusInternalServerError)
		return
	}
	// remove from sessions db
	delete(dbSessions, cookie.Value)

	// expire the cookie.
	cookie.Value = ""
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func signUp(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		return
	}

	if req.Method == http.MethodPost {
		// need username, pwd,, first, last, role
		un := req.FormValue("username")
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		}

		p := req.FormValue("pwd")
		f := req.FormValue("first")
		l := req.FormValue("last")
		r := req.FormValue("role")

		pwd, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		u := user{un, f, l, pwd, r}
		// we need to be careful here to ensure we dont end up with two sessions for same user if have two browsers.
		// We can have a race conditon here wherein we can end up with multiple sessions for the same user.
		dbUsers[un] = u // in real world this api would be "insert only it not exists else return error". On error we will not create session since user would already exist and may have a session
		sid := createSession(un)
		cookie := &http.Cookie{
			Name:   sidName,
			Value:  sid,
			MaxAge: maxAge,
		}

		http.SetCookie(w, cookie)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		log.Println("User created: ", un, "Session ", sid, " Created.", "Active sessions: ", len(dbSessions))
		return
	}

	if err := tpl.ExecuteTemplate(w, "signup.gohtml", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go cleanSessions(ctx)
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/signUp", signUp)
	http.HandleFunc("/userExists", userExists)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
