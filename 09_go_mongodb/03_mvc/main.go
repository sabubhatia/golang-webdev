package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sabubhatia/golang-webdev/09_go_mongodb/03_mvc/controllers"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Panic(err)
	}
	return s
}

func main() {
	mux := httprouter.New()
	s := getSession()
	defer s.Close()
	uc := controllers.NewUserController(s, "go-web-dev-db")
	mux.GET("/", uc.Index)
	mux.GET("/user/:id", uc.GetUser)
	mux.POST("/user", uc.CreateUser)
	mux.DELETE("/user/:id", uc.DeleteUser)
	log.Println("ListendAndServe()...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
