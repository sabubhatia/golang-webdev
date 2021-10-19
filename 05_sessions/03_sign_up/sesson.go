package main

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func newsSessionCookie() (*http.Cookie, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	cookie := &http.Cookie{
		Name:     "sid",
		Value:    uid.String(),
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		MaxAge:   SessionLength,
	}

	return cookie, nil
}
func getUser(w http.ResponseWriter, req *http.Request) user {
	// 1. Cookie must exist.
	// 2. sid->un exists.
	// 3. un->user exists.
	var u user
	cookie, err := req.Cookie("sid")
	if err != nil {
		return u
	}

	if s, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[s.un]
		s.lastActivity = time.Now()
		dbSessions[cookie.Value] = s
		cookie.MaxAge = SessionLength
		// set the cookie here.
		http.SetCookie(w, cookie)
	}

	return u
}

func alreadyLogedIn(w http.ResponseWriter, req *http.Request) bool {
	// If have a cookie && have  sid->un mapping and a un->user mapping then logged in
	cookie, err := req.Cookie("sid")
	if err != nil {
		return false
	}
	f := func() {
		cookie.MaxAge = SessionLength
		http.SetCookie(w, cookie)
	}
	// does sid->un exist in session db.
	if s, ok := dbSessions[cookie.Value]; ok {
		// does un->user exist in the user dataabse.
		if _, ok := dbUsers[s.un]; !ok {
			// Will  only get here is the user has been deleted from the db.
			return false
		}
		s.lastActivity = time.Now()
		dbSessions[cookie.Value] = s
		f() // Cookie and session last activity are at parity.
	} else {
		return false
	}

	return true
}

func cleanDBSessions() int {

	cnt := 0
	for k, v := range dbSessions {
		if time.Since(v.lastActivity) >= (time.Second * time.Duration(SessionLength)) {
			// Elapsed stay alive time so delete this.
			delete(dbSessions, k)
			cnt++
		}
	}

	return cnt
}
