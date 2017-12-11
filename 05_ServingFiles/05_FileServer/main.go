package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("."))) //FileServer returns an handler
	http.HandleFunc("/dog", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		io.WriteString(w, `<img src="tadm.jpg">`)
	})

	http.ListenAndServe(":8080", nil)
}
