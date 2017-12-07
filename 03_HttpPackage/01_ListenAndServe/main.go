package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (n hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) { //hotdog implements handler interface
	fmt.Println(w, "Any code you want")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
