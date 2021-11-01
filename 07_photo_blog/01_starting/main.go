package main

import (
	"crypto/sha256"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var (
	tpl   *template.Template
	port  = ":8080"
	delim = "|"
)

func init() {
	if len(os.Args) < 2 {
		log.Fatal("Expected at least 2 args. Got: ", len(os.Args))
	}

	tpl = template.Must(template.ParseGlob(os.Args[1]))
	log.Println(tpl.DefinedTemplates())

	if len(os.Args) > 2 {
		port = os.Args[2]
	}
}

func cookieValue(value, substr string) (string, error) {
	// Generate a cookie value that has the format:
	// uuid|substr|substr..
	if len(value) <= 0 {
		// New cookie value so generate a uuid.
		uuid, err := uuid.NewV4()
		if err != nil {
			return "", err
		}

		value += uuid.String()
	}

	if len(substr) <= 0 {
		return value, nil
	}

	if strings.Contains(value, substr) {
		return value, nil // already exists
	}

	// Have a non zero string in value and a non zero substr that is not in value. SO add substr
	return value + delim + substr, nil
}

func getCookie(req *http.Request) (*http.Cookie, error) {
	cookie, err := req.Cookie("sID")
	if err != nil && err != http.ErrNoCookie {
		return nil, err
	}
	return cookie, nil
}

func setCookie(w http.ResponseWriter, req *http.Request, substr string) *http.Cookie {
	cookie, err := getCookie(req)
	if err != nil {
		handleErr(w, err)
		return nil
	}

	var val string
	if cookie != nil {
		val = cookie.Value
	}
	val, err = cookieValue(val, substr)
	if err != nil {
		handleErr(w, err)
		return nil
	}

	cookie = &http.Cookie{
		Name:     "sID",
		Value:    val,
		Path:     "/",
		MaxAge:   30,
		Secure:   false, // not using https currently
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	return cookie
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie := setCookie(w, req, "")
	if cookie == nil {
		return
	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", cookie)
	handleErr(w, err)
}

func addFile(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("fname")
		if err != nil {
			http.Redirect(w, req, "/result?val="+err.Error(), http.StatusSeeOther)
			return
		}
		defer mf.Close()

		// Name of file I will store will be hash(file contents) plus passed in extension.
		// Extract file extension from file header. Next read file data and then hash it.

		ext := path.Ext(fh.Filename)
		if len(ext) <= 0 {
			http.Redirect(w, req, "/result?val= Cannot decipher extension from file name: "+fh.Filename, http.StatusSeeOther)
			return
		}

		h := sha256.New()
		// Now write the file data to the hash.
		_, err = io.Copy(h, mf)
		if err != nil {
			http.Redirect(w, req, "/result?val= Hash failed: "+err.Error(), http.StatusSeeOther)
			return
		}

		// File name = hash value + extenstion
		fname := fmt.Sprintf("%x%s", h.Sum(nil), ext)

		// File will be created in "current working directory/public/pics and will have the name fn"
		wd, err := os.Getwd()
		if err != nil {
			http.Redirect(w, req, "/result?val= Unable to get current working directory: "+err.Error(), http.StatusSeeOther)
			return
		}
		fn := filepath.Join(wd, "public", "pics", fname)
		if len(fn) <= 0 {
			// This path in code is really not possible given all the checks done above. But good hygeine 
			// to test for all error conditions.
			http.Redirect(w, req, "/result?val= Created file name is empty", http.StatusSeeOther)
			return
		}

		// Now create the file and write the data to it. We had already read the data for the hash.
		// So we need to seek back to the start of the mf reader.

		f, err := os.Create(fn)
		if err != nil {
			http.Redirect(w, req, "/result?val= Unable to create file: " + fn + "," + err.Error(), http.StatusSeeOther)
			return
		}
		del := false // Default is it is all ok. So dont delete fiel in "f"
		defer func() {
			f.Close()
			if del {
				// remove this file.
				os.Remove(fn)
			}
		}()

		_, err = mf.Seek(0, io.SeekStart)
		if err != nil {
			http.Redirect(w, req, "/result?val= Unable to seek multi file reader to start: "+err.Error(), http.StatusSeeOther)
			// delete the created file.
			del  = true
			return
		}

		_, err = io.Copy(f, mf)
		if err != nil {
			http.Redirect(w, req, "/result?val= Unable to copy file: "+err.Error(), http.StatusSeeOther)
			// delete the created file.
			del  = true
			return		
		}

		// Cookie only gets set with this file if we can create it successfully. Not we only use fname not the
		// file name with a full qualified path
		cookie := setCookie(w, req, fname)
		if cookie == nil {
			http.Redirect(w, req, "/result?val=Unable to set cookie", http.StatusSeeOther)
			del = true // deleting file since cookie cant be set so the file will be asssumed to not exist.
			return
		}
		http.Redirect(w, req, "/result?val=Cookie set too: "+cookie.Value, http.StatusSeeOther)
		return
	}

	cookie := setCookie(w, req, "")
	if cookie == nil {
		return
	}
	xs := strings.Split(cookie.Value, delim)
	err := tpl.ExecuteTemplate(w, "file.gohtml", xs[1:]) // Skip the uuid and only show the pics
	handleErr(w, err)
}

func result(w http.ResponseWriter, req *http.Request) {
	val := req.FormValue("val")
	err := tpl.ExecuteTemplate(w, "result.gohtml", val)
	handleErr(w, err)
}

func handleErr(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/addFile", addFile)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	http.HandleFunc("/result", result)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("ListenAndServe() on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
