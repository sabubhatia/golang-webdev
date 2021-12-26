package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sabubhatia/golang-webdev/11_postgres/03_web_app/model"
)

type BookController struct {
	db *sql.DB
}

func NewBookController(db *sql.DB) *BookController {
	return must(db)
}

func must(db *sql.DB) *BookController {
	if db == nil {
		panic("Passed in db pointer cannot be nil")
	}

	return &BookController{db}
}

func (bc BookController) GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := bc.db.Query("select * from books")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	bsx := make([]model.Book, 0)
	for rows.Next() {
		b := model.Book{}
		if err := rows.Scan(&b.Isbn, &b.Title, &b.Author, &b.Price); err != nil {
			break
		}
		bsx = append(bsx, b)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bsx); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
