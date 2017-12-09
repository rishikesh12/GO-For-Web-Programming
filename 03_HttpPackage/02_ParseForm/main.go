package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

var tpl *template.Template

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, r.Form) //Data is parsed from the form using ParseForm which returns a map
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

//Form values is of map type which has a string as a key and a slice of sting as value
//so the values entered in the form will be stored in the map
