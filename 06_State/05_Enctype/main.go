package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))

}
func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	xs := make([]byte, r.ContentLength)
	r.Body.Read(xs)

	err := tpl.Execute(w, string(xs))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

//enctype stands for encoding type
//refer index.gohtml for different enctype options
