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
	io.WriteString(w, "Cat")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", d) //catches anything down the line after dog
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux) //servemux implements servehttp method hence it is a handler
}
