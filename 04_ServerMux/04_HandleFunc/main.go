package main

import (
	"io"
	"net/http"
)

type hotdog int

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Cat")
}

func main() {
	http.HandleFunc("/dog/", d) //handlefunc takes a func(ResponseWriter,*Request)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil) //we are using the default ServeMux
}

//HandlerFunc is different from HandleFunc (see the documentation)
//HandlerFunc is basically an helper function  which is used to convert an handler
//type handlerFunc implements ServeHTTP
//http.Handle("/dog",http.HandlerFunc(d))
