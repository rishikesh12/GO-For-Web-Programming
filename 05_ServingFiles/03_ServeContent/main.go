package main

import (
	"io"
	"log"
	"net/http"
	"os"
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
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("tadm.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}

//ServeContent is rarely used
