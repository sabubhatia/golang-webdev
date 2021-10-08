package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var tpl *template.Template

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("Expeted at least two args. Got: %d", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func foo(w http.ResponseWriter, req *http.Request) {

	var s string
	var first string
	if req.Method == http.MethodPost {
		f, fh, err := req.FormFile("file")
		if err != nil {
			HandleError(w, err)
			return
		}
		defer f.Close()
		log.Println("File Header:\n", fh, "\n", "File:\n", f)
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			HandleError(w, err)
			return
		}

		s = string(bs)
		first = req.FormValue("first")

		// write the file to the server.

		dst, err := os.Create(filepath.Join("./tmp", first+".txt"))
		if err != nil {
			HandleError(w, err)
			return
		}
		defer dst.Close()
		_, err = dst.Write(bs)
		if err != nil {
			HandleError(w, err)
			return
		}

		log.Println(req.Form)
		log.Println(req.MultipartForm)
	}

	d := struct {
		F     string
		First string
	}{
		s,
		first,
	}
	err := tpl.ExecuteTemplate(w, "ReadFile", d)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
