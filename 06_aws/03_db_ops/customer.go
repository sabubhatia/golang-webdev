package main

import (
	"io"
	"log"
	"net/http"
)

const (
	custTable = "customer"
)

func custTableExists() (bool, error) {
	r, err := db.Query("SELECT count(*) FROM information_schema.tables where TABLE_NAME=" + `"` + custTable + `"` + ";")
	if err != nil {
		return false, err
	}
	defer r.Close()

	for r.Next() {
		var cnt int
		if err := r.Scan(&cnt); err != nil {
			return false, err
		}
		if cnt <= 0 {
			return false, nil
		}
	}

	return true, nil
}

func createCust(w http.ResponseWriter, req *http.Request) {
	ex, err := custTableExists()
	log.Println(ex, err)
	if err != nil {
		handleErr(w, err)
		return
	}
	if ex {
		http.Redirect(w, req, "/result?val=Table customer already exsits", http.StatusSeeOther)
		return
	}

	s, err := db.Prepare(`CREATE table ` + custTable + `(
		cID INT NOT NULL AUTO_INCREMENT,
		cName VARCHAR(256) NOT NULL,
		cAddr VARCHAR(512) NOT NULL,
		cEmail VARCHAR(256) NOT NULL,
		PRIMARY KEY (cID),
		UNIQUE INDEX cID_UNIQUE (cID ASC) VISIBLE);`)
	if err != nil {
		handleErr(w, err)
		return
	}
	defer s.Close()
	if _, err := s.Exec(); err != nil {
		handleErr(w, err)
		return
	}
	http.Redirect(w, req, "/result?val=Table customer successfully created", http.StatusSeeOther)
}

func dropCust(w http.ResponseWriter, req *http.Request) {
	ex, err := custTableExists()
	if err != nil {
		handleErr(w, err)
		return
	}
	if !ex {
		http.Redirect(w, req, "/result?val=Table customer does not exist", http.StatusSeeOther)
		return
	}
	s, err := db.Prepare("DROP table " + custTable)
	if err != nil {
		handleErr(w, err)
		return
	}
	defer s.Close()

	if _, err := s.Exec(); err != nil {
		handleErr(w, err)
		return
	}

	http.Redirect(w, req, "/result?val=Table customer has been successully deleted", http.StatusSeeOther)
}

func addCust(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "addCust() <br>")
	io.WriteString(w, "<a href=/> Home </a> <br>")
}

func getCust(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "getCust() <br>")
	io.WriteString(w, "<a href=/> Home </a> <br>")
}

func searchCust(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "searchCust() <br>")
	io.WriteString(w, "<a href=/> Home </a> <br>")
}
