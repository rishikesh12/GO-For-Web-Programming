package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Foo Method:", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bar Method:", r.Method)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Barred Mehod:", r.Method)
	tpl.Execute(w, nil)

}

//the submitted form will be handled by bar(which redirects it to foo)
//In temporary redirect the post method used by the form remains the same
