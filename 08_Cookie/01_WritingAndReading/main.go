package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "new-cookie",
		Value: "some-value",
	}) //setting a cookie
	fmt.Fprintln(w, "cookie written")
	fmt.Fprintln(w, "Go to dev-tools/application/cookies")
}

func read(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("new-cookie") //checking whether a cookie with the name new-cookie exists
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintln(w, "Your cookie", c)

}

//cookies are domain specific
//cookie is a struct in the http package
