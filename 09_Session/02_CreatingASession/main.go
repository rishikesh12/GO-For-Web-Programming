package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

type user struct {
	Uname    string
	First    string
	Last     string
	LoggedIn bool
}

var dbsession = map[string]string{} // key uuid and value uname
var db = map[string]user{}          // Map with uname as key and the user as value

func init() {
	//fmt.Println(os.Getwd())
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	//http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	//Creating a cookie if it doesn't exist
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
		http.SetCookie(w, c)
	}
	//Process the form submission
	var u user
	if r.Method == http.MethodPost {
		uname := r.FormValue("uname")
		f := r.FormValue("fname")
		l := r.FormValue("lname")
		li := r.FormValue("loggedin") == "on" //on is a keyword for checkbox
		u = user{uname, f, l, li}
		dbsession[c.Value] = uname
		db[uname] = u
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		log.Println(err)
	}

}

func bar(w http.ResponseWriter, r *http.Request) {
	//get the cookie
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	uname, ok := dbsession[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	err = tpl.ExecuteTemplate(w, "bar.gohtml", db[uname])
	if err != nil {
		log.Println(err)
	}

}

//comma ok idiom : A map always returns a zero values for any value that you might have not put in the Map
// so we use ok to check whether we have put a particluar value in the map
