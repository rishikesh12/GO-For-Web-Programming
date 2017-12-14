package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method inside foo:", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method inside bar", r.Method)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//When using StatusPermanentRedirect the browser remembers the redirect Location and
//from the next time the new handler will handle the request
