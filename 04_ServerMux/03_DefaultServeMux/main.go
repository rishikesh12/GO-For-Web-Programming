package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "cat")
}

func main() {
	var d hotdog
	var c hotcat
	http.Handle("/dog", d)
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil) //nil signifies that we will be using DefaultServeMux
}
