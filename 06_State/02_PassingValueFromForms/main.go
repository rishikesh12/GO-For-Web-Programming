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

	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `
    <form method="get">
    <input type="text" name="q">
    <input type="submit">
    </form>
    </br> `+v)
}

//get passes the values through the url
//post passes the value through the body
