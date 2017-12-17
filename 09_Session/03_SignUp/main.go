package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

type user struct {
	Uname string
	Fname string
	Pass  string
}

var dbsession = map[string]string{} //key uuid and value username
var dbuser = map[string]user{}      //key username and value user
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getuser(w, r)
	err := tpl.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		log.Println(err)
	}

}

func signup(w http.ResponseWriter, r *http.Request) {

	check := alreadyregistered(r)
	if check {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	//get form database

	if r.Method == http.MethodPost {
		uname := r.FormValue("uname")
		fname := r.FormValue("fname")
		pass := r.FormValue("pass")

		//check if username is taken
		if _, ok := dbuser[uname]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
		}

		//create Session
		id := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
		http.SetCookie(w, c)

		u := user{uname, fname, pass}
		//register the new user in the database
		dbsession[c.Value] = uname
		dbuser[uname] = u

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}
	err := tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	if err != nil {
		log.Println(err)
	}

}

func bar(w http.ResponseWriter, r *http.Request) {
	//check if user exists
	check := alreadyregistered(r)
	if !check {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}
	u := getuser(w, r)
	err := tpl.ExecuteTemplate(w, "bar.gohtml", u)
	if err != nil {
		log.Println(err)
	}

}
