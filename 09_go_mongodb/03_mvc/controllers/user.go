package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/sabubhatia/golang-webdev/09_go_mongodb/03_mvc/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	s *mgo.Session
	db string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func NewUserController(s *mgo.Session, db string) *UserController {
	return &UserController{s, db}
}

func (Uc UserController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	i := 1000
	if err := tpl.ExecuteTemplate(w, "index.gohtml", i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(tpl.DefinedTemplates())
}

func (Uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "not a valid Hex ID")
		return
	}

	// convert hex id back to obejct id.
	oid := bson.ObjectIdHex(id)

	u := models.User{}
	// Now read from mongodb collection the user identified by this object id.
	if err := Uc.s.DB(Uc.db).C(models.Collection).FindId(oid).One(&u); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func (Uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get the bson id for mongodb
	u.Id = bson.NewObjectId()

	if err := Uc.s.DB(Uc.db).C(models.Collection).Insert(u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}
	// Marshall it back.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func (Uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// get object ID from hex
	oid := bson.ObjectIdHex(id)
	if err := Uc.s.DB(Uc.db).C(models.Collection).RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, id + ", Has been deleted !")
}
