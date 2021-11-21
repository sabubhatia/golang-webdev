package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type userid string

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(req.Context(), userid("uid"), 777)	
	fmt.Fprintln(w, "UID: ", dbRead(ctx))
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(req.Context(), userid("sid"), 1000)
	fmt.Fprintln(w, req.Context())
	fmt.Fprintln(w, doStuff(ctx))
}

func dbRead(ctx context.Context) int {
	uid, ok := ctx.Value(userid("uid")).(int)
	if !ok {
		log.Println("int expected: ", uid)
		return -777
	}

	return uid
}

func doStuff(ctx context.Context) int {
	sid, ok := ctx.Value(userid("sid")).(int)
	if !ok {
		log.Println("int expected")
		return -1000
	}

	return sid
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
