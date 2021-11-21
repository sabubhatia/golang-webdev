package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type stringid string

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(req.Context(), stringid("uid"), 777)
	v, err := dbAccess(ctx)
	if err != nil {
		log.Println(err.Error())
		log.Println("Value is: ", v)
		time.Sleep(15 * time.Second)
		log.Println("Woke up..")
		return
	}

	log.Println("No err. value is: ", v)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, req.Context())
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	chi := make(chan int)
	// start an absrudly slow db read operation.
	go func() {
		defer close(chi)
		uid, ok := ctx.Value(stringid("uid")).(int)
		if !ok {
			log.Fatal("Logic error. Expected int")
		}
		time.Sleep(time.Second * 10) // emulate a slow running query
		if ctx.Err() != nil {
			// too slow context is canceled or timed out
			return
		}

		log.Println("Writing  to chi")
		chi <- uid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case v := <-chi:
		return v, nil
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
