package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sabubhatia/golang-webdev/09_go_mongodb/02_json/models"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	i := 1000
	if err := tpl.ExecuteTemplate(w, "index.gohtml", i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(tpl.DefinedTemplates())
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Sabu",
		Gender: "Male",
		Age:    52,
		Id:     p.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(uj))
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// do something to show it is created.
	u.Id = u.Id + ":OK"

	// Marshall is back.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, id)
}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/user/:id", getUser)
	mux.POST("/user", createUser)
	mux.DELETE("/user/:id", deleteUser)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
