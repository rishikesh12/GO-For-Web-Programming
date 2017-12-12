package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/tadm.jpg", dogPic)
	http.ListenAndServe(":8081", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="/tadm.jpg">
	`) //this path redirects to dogPic func where iocopy is used
}ud

func dogPic(w http.ResponseWriter, req *http.Request) {

	http.ServeFile(w, req, "tadm.jpg")
}
