package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver       = "mysql"
	dataSourceName = "admin:password@tcp(database-1.cauehzhs4qsn.ap-southeast-1.rds.amazonaws.com:3306)/test01?charset=utf8"
)

var (
	port = ":8080"
	tpl  *template.Template
	db   *sql.DB
	mach *machine
)

func init() {
	log.Println("init()...")
	// args: 1 = templats, 2 port. 2 is optional.
	if len(os.Args) < 2 {
		log.Fatal("Expected at least 2 arg. Got: ", len(os.Args))
	}
	if len(os.Args) > 2 {
		// port number
		port = os.Args[2]
	}

	var err error
	if mach, err = getMachine(); err != nil {
		log.Fatal("Unable to get machine details: " + err.Error())
	}

	// load the templates
	tpl = template.Must(template.ParseGlob(os.Args[1]))
	log.Println(tpl.DefinedTemplates())

	// get a connection to the DB and make sure it is reachable.
	db, err = sql.Open(dbDriver, dataSourceName)
	if err != nil {
		log.Fatal("Unable to establish connection to DB. ", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Unable to ping DB. ", err.Error())
	}
	log.Println("init() done...InstanceID: ", mach.InstanceID)

}

func done() {
	db.Close()
}
func index(w http.ResponseWriter, req *http.Request) {
	d := struct {
		M *machine
	} {
		M: mach,
	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", d)
	handleErr(w, err)
}

func result(w http.ResponseWriter, req *http.Request) {
	d := struct {
		V string
		M *machine
	} {
		V: req.FormValue("val"),
		M: mach,
	}
	err := tpl.ExecuteTemplate(w, "result.gohtml", d)
	handleErr(w, err)
}

func instanceShow(w http.ResponseWriter, request *http.Request) {
	err := tpl.ExecuteTemplate(w, "instance.gohtml", mach)
	handleErr(w, err)
}

func insert(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		nm := req.FormValue("name")
		// if empty return error
		if nm == "" {
			http.Redirect(w, req, "/result?val=Empty name not permitted", http.StatusSeeOther)
			return
		}
		// if exists return error
		r, err := db.Query(`SELECT count(aName) from amigos where aName = ` + `"` + nm + `"` + `;`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Close()
		for r.Next() {
			// Get the value if > 0 then exists.
			var c int
			if err := r.Scan(&c); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if c > 0 {
				http.Redirect(w, req, "/result?val=Name already exists", http.StatusSeeOther)
				return
			}
		}
		// prepare an insert statement then execute the statement.
		s, err := db.Prepare("Insert into amigos (aName) Values " + `("` + nm + `");`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer s.Close()
		var res sql.Result
		res, err = s.Exec()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if i, _ := res.RowsAffected(); i <= 0 {
			// At this point should not really happen. So throwing error
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Done do redirect to a suceess
		cnt, _ := res.RowsAffected()
		id, _ := res.LastInsertId()
		msg := fmt.Sprintf("Succesffuly added %d row(s). Added row ID is: %d", cnt, id)
		http.Redirect(w, req, "/result?val="+msg, http.StatusSeeOther)
		return
	}

	d := struct {
		M *machine
	} {
		M: mach,
	}
	err := tpl.ExecuteTemplate(w, "insert.gohtml", d)
	handleErr(w, err)
}

func delete(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		nm := req.FormValue("name")
		if nm == "" {
			http.Redirect(w, req, "/result?val=Empty string is not a valid amigo name", http.StatusSeeOther)
			return
		}

		stmt, err := db.Prepare("DELETE from amigos where aName = ?")
		if err != nil {
			handleErr(w, err)
			return
		}
		defer stmt.Close()
		res, err := stmt.Exec(nm)
		if err != nil {
			handleErr(w, err)
			return
		}

		if cnt, _ := res.RowsAffected(); cnt <= 0 {
			http.Redirect(w, req, "/result?val=Nothing deleted from database", http.StatusSeeOther)
			return
		}

		http.Redirect(w, req, "/result?val=Amigo has been successfully deleted", http.StatusSeeOther)
		return
	}

	d := struct {
		M *machine
	} {
		M: mach,
	}
	err := tpl.ExecuteTemplate(w, "delete.gohtml", d)
	handleErr(w, err)
}

func read(w http.ResponseWriter, req *http.Request) {
	q := `SELECT aName from amigos`
	r, err := db.Query(q)
	if err != nil {
		handleErr(w, err)
		return
	}
	defer r.Close()

	sx := []string{}
	for r.Next() {
		var nm string
		if err := r.Scan(&nm); err != nil {
			break
		}
		sx = append(sx, nm)
	}

	if err := r.Err(); err != nil {
		handleErr(w, err)
		return
	}

	d := struct {
		Sx []string
		Q  string
		M *machine
	}{
		Sx: sx[0:],
		Q:  q,
		M: mach,
	}
	err = tpl.ExecuteTemplate(w, "select.gohtml", d)
	handleErr(w, err)
}

func handleErr(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func ping(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/http")
	io.WriteString(w, "Ping..." + mach.InstanceID)
}

func main() {
	defer done() // tear down anything that was allocated in init
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/read", read)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/createCust", createCust)
	http.HandleFunc("/dropCust", dropCust)
	http.HandleFunc("/addCust", addCust)
	http.HandleFunc("/getCust", getCust)
	http.HandleFunc("/searchCust", searchCust)
	http.HandleFunc("/instance", instanceShow)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/result", result)
	log.Fatal(http.ListenAndServe(port, nil))
}
