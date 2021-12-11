package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	maxAge                  = 30 //seconds
	sidName                 = "SID"
	cleanTime time.Duration = 30 //seconds
)

var (
	lastCleaned = time.Now()
)

func createSession(u string) string {
	s := session{
		Un:           u,
		LastActivity: time.Now(),
	}

	sid, err := uuid.NewV4()
	if err != nil {
		log.Fatal("UUID generation failure: " + err.Error())
	}
	dbSessions[sid.String()] = s

	return sid.String()
}

func getUser(w http.ResponseWriter, req *http.Request) user {

	cookie, err := req.Cookie(sidName)
	if err == nil && alreadyLoggedIn(w, req) {
		s, ok := dbSessions[cookie.Value]
		if ok {
			return dbUsers[s.Un]
		}
	}

	// Only way we are here is:
	// (1) No cookie exists.
	// OR
	// (2) if between map look up and alreadylogged() there was a time slice of greater
	// than cleanTime that caused the session to be inactive and cleaned out.
	// OR
	// (3) getUser() was called before.
	// Handle this by sending out the empty new user
	// Either way there is no session started.
	sid, err := uuid.NewV4()
	if err != nil {
		log.Fatal("UUID generation failure: " + err.Error())
	}
	cookie = &http.Cookie{
		Name:   sidName,
		Value:  sid.String(),
		MaxAge: maxAge,
	}

	http.SetCookie(w, cookie)
	return user{} // Empty value
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie(sidName)
	if err == nil {
		s, ok := dbSessions[cookie.Value]
		if ok {
			// update the last activity stamp in the db sessions
			s.LastActivity = time.Now()
			dbSessions[cookie.Value] = s
			// update the cookie.
			cookie.MaxAge = maxAge
			http.SetCookie(w, cookie)
			return true
		}
	}

	// No have a cookie or by the time we got to it another channel logged us out or was never logged in
	return false
}

func cleanSessions(ctx context.Context) {

	ticker := time.NewTicker(cleanTime * time.Second)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("Tearing down cleanSessions()..")
			return
		case <-ticker.C:
			cleanSessionsDB()
			log.Println("Cleaned at: ", fmtTime(lastCleaned))
		}
	}
}

func fmtTime(t time.Time) string {
	return t.Format(time.StampMilli)
}

func cleanSessionsDB() {
	for k, v := range dbSessions {
		if time.Since(v.LastActivity) > (time.Second * cleanTime) {
			log.Println("Cleaning SID: ", k, time.Since(v.LastActivity), time.Second * cleanTime )
			delete(dbSessions, k)
		}
	}
	lastCleaned = time.Now()
}

func showSessions(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	defer func() {
		// Write the footers.
		fmt.Fprintf(w, `<a href="/login"> Login </a> <br>`)
		fmt.Fprintf(w, `<a href="/signUp"> Sign Up </a> <br>`)
	}()
	if len(dbSessions) <= 0 {
		io.WriteString(w, "No active sessions found. <br>")
		return
	}
	io.WriteString(w, "<ul>")
	for k, v := range dbSessions {
		fmt.Fprintf(w, "<li> %s: %s, %s </li>", k, v.Un, fmtTime(v.LastActivity))
	}
	io.WriteString(w, "</ul>")
}
