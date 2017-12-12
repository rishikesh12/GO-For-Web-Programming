package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/dog", dog)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintln(w, "Check your terminal")
}

//if the browser asks for a favicon the program will writes a 404 PageNotFound message
