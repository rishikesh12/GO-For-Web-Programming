package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (n hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Krishna", "the best in the world") //Header is a map
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want</h1>")
}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
