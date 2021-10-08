package main

import (
	"net/http"

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
		Secure:   true,
		HttpOnly: true,
	}

	return cookie, nil
}
func getUser(req *http.Request) user {
	// 1. Cookie must exist.
	// 2. sid->un exists.
	// 3. un->user exists.
	var u user
	cookie, err := req.Cookie("sid")
	if err != nil {
		return u
	}

	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	return u
}

func alreadyLogedIn(req *http.Request) bool {
	// If have a cookie && have  sid->un mapping and a un->user mapping then logged in
	cookie, err := req.Cookie("sid")
	if err != nil {
		return false
	}

	// does sid->un exist in session db.
	if un, ok := dbSessions[cookie.Value]; ok {
		// does un->user exist in the user dataabse.
		if _, ok := dbUsers[un]; !ok {
			return false
		}
	} else {
		return false
	}

	return true
}
