package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"runtime"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string
	First    string
	Last     string
	Pwd      []byte
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

type env struct {
	OS string
	Arch string
	Ncpu int
}
var (
	dbUsers     = map[string]user{}    // un->user
	dbSessions  = map[string]session{} // sid->session
	tpl         *template.Template
	lastCleaned = time.Now()
	chanClean   = make(chan struct{})
	runEnv env = env{OS: runtime.GOOS, Arch:runtime.GOARCH, Ncpu: runtime.NumCPU(),}
)

const (
	RoleAdmin         = "admin"
	RoleRead          = "read"
	RoleRW            = "rdwrt"
	RoleWrite         = "write"
	SessionLength int = 30 // seconds.
)

func init() {
	// load users data
	pwd, err := bcrypt.GenerateFromPassword([]byte("abcdef"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	dbUsers["sabusingh.bhatia@gmail.com"] = user{
		Username: "sabusingh.bhatia@gmail.com",
		First:    "Sabu",
		Last:     "Bhatia",
		Pwd:      pwd,
		Role:     "admin",
	}
	if len(os.Args) < 2 {
		log.Fatal("Expected at least 2 args. Got: ", len(os.Args))
	}


	log.Println("Args: ", os.Args[1])
	tpl = template.Must(template.ParseGlob(os.Args[1]))
	log.Println(tpl.DefinedTemplates())
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	d := struct {
		U user
		E env
	} {
		U: u,
		E: runEnv,
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	if !alreadyLogedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	u := getUser(w, req)
	if u.Role != RoleRW {
		http.Error(w, "You dont have sufficient priviliges to enter the bar.", http.StatusForbidden)
		return
	}
	err := tpl.ExecuteTemplate(w, "bar.gohtml", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLogedIn(w, req) {
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
		hp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		r := req.FormValue("role")
		cookie, err := newsSessionCookie()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = session{un, time.Now()}
		u = user{un, f, l, hp, r}
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
	log.Println("login()..")
	if alreadyLogedIn(w, req) {
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

		log.Println("Processing login()..", un)
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
		dbSessions[cookie.Value] = session{un, time.Now()}
		log.Println("login() processed")

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
	if !alreadyLogedIn(w, req) {
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
		// Pehaps it may be better to throw a fatal error here. Basically the observability api's
		// you may have come in handy here to decide when to start cleaning. Remember in the real world this can while the 
		// cleaning is runnign degrade performance  
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
	u := getUser(w, req)
	if u.Role != RoleAdmin {
		http.Error(w, "You must have admin priviliges to see the dump", http.StatusForbidden)
		return
	}
	d := struct {
		DBSessions map[string]session
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // On exit clean up.
	go clean(ctx, chanClean)
	go tick(ctx, chanClean)

	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/dump", dump)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":80", nil))
}

func tick(ctx context.Context, out chan<- struct{}) {
	// using a timer here. But in the real world shoudl be perhaps based on factors such as
	// current system load, or the size of the dbsessions table or a certain window in which house 
	// keeping can be done. It can be a combination of these. Your observability api's come in handy here.
	ticker := time.NewTicker(time.Second* time.Duration(SessionLength * 4))
	defer func() {
		log.Println("Stopping ticker()..")
		ticker.Stop()
	}()

	log.Println("tick() started...")
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			out <- struct{}{}
		}
	}
}

func clean(ctx context.Context, in <-chan struct{}) {
	log.Println("clean() started...")
	for {
		select {
		case <-ctx.Done():
			log.Println("Context done:", ctx.Err())
			return
		case <-in:
			log.Println("Cleaning sessions..")
			if time.Since(lastCleaned) >= time.Second*time.Duration(SessionLength) {
				log.Println("#Sessionss cleaned: ", cleanDBSessions())
				lastCleaned = time.Now()
			}
		}
	}
}

