package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("qe")
	io.WriteString(w, "The passed parameter is :"+v)
}

//We can use FormValue to access the passed paramters
