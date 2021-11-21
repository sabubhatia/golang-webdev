package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	tpl *template.Template
)

func init() {
	if len(os.Args) < 2 {
		log.Fatal("Expected at least 2 Args. Got: ", len(os.Args))
	}

	tpl = template.Must(template.ParseGlob(os.Args[1]))
	log.Println(tpl.DefinedTemplates())
}

func getCode(s string) string {
	bx := bytes.Join([][]byte{[]byte(s), []byte("My super secret key")}, []byte{})
	h := hmac.New(sha256.New, bx)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func cmp(s, h string) bool {
	log.Println(s, h)
	return hmac.Equal([]byte(getCode(s)), []byte(h))
}

func foo(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("sID")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "sID",
			Value: "",
		}
	}

	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		if e == "" {
			http.Redirect(w, req, "/result=Email cannot be empty", http.StatusSeeOther)
			return
		}
		cookie.Value = e + "|" + getCode(e)
	}

	http.SetCookie(w, cookie)
	err = tpl.ExecuteTemplate(w, "index.gohtml", cookie)
	handleErr(w, err)
}

func authenticate(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("sID")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if cookie.Value == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	sx := strings.Split(cookie.Value, "|")
	if len(sx) != 2 {
		http.Redirect(w, req, "/result?val=Cookie split string value "+strconv.Itoa(len(sx))+" not valid", http.StatusSeeOther)
	}

	err = tpl.ExecuteTemplate(w, "auth.gohtml", cmp(sx[0], sx[1]))
	handleErr(w, err)
}

func handleErr(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/authenticate", authenticate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
