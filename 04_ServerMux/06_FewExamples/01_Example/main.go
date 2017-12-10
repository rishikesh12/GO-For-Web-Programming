package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Index")
	})

	http.HandleFunc("/dog", d)

	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog")
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("rishi.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(w, "Rishi")
	if err != nil {
		log.Fatalln(err)
	}

}
