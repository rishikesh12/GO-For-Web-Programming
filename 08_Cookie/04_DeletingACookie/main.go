package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">set a cookie</a><h1>`)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "my-value",
	})
}

func read(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintln(w, cookie)

}

func expire(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}
	cookie.MaxAge = -1 //Delete cookie
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// MaxAge indicates the cookie's life in seconds
