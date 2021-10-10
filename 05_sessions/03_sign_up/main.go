package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string
	First    string
	Last     string
	Pwd      []byte
}

var (
	dbUsers    = map[string]user{}   // un->user
	dbSessions = map[string]string{} // sid->un
	tpl        *template.Template
)

func init() {
	// load users data
	pwd, err := bcrypt.GenerateFromPassword([]byte("abcdef"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	dbUsers["sabusingh.bhatia@gmail.com"] = user{
		Username: "sabusingh.bhatia@gmail.com",
		First: "Sabu",
		Last: "Bhatia",
		Pwd: pwd,
	}
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
		hp, err  := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}

		cookie, err := newsSessionCookie()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = un
		u = user{un, f, l, hp}
		dbUsers[un] = u

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	err := tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLogedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		// process login

		// does the user have an account?
		un := req.FormValue("username")
	    u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Unrecognised username or password", http.StatusForbidden)
			return
		}

		// compare the encrypted password to the passed in password
		pwd := req.FormValue("password") 
		err := bcrypt.CompareHashAndPassword(u.Pwd, []byte(pwd))
		if err != nil {
			http.Error(w, "Unrecognised username or password", http.StatusForbidden)	
			return 
		}

		// create a session id sid.
		cookie, err := newsSessionCookie()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}

		// Set the cookie to the session id
		http.SetCookie(w, cookie)

		// update session table. map sid to user name.
		dbSessions[cookie.Value] = un

		// Logged on. Go to the index.
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLogedIn(req) {
		log.Println("Not logged in...")
		// Not logged in so redirect somewhere else.
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return 
	}

	log.Println("Logout...")
	cookie, err := req.Cookie("sid")
	if err != nil {
		// This shouldnt really be here given we are per above logged in so a cookie must exist.
		// We are still handling the error. But if we are here something is really wrong.
		// Pehaps it may be better to throw a fatal error here.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// expire the cookie.
	cookie.MaxAge = -1

	// delete entry from session map
	delete(dbSessions, cookie.Value)

	// set the cookie.
	http.SetCookie(w, cookie)

	http.Redirect(w, req, "/login", http.StatusSeeOther)
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
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
