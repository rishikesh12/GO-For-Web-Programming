package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "new-cookie",
		Value: "new-value",
	})
}

func read(w http.ResponseWriter, r *http.Request) {

	c1, err := r.Cookie("new-cookie")
	if err != nil {
		log.Fatal("Cookie not found")
	} else {
		fmt.Fprintln(w, c1)
	}
	c2, err := r.Cookie("old-cookie")
	if err != nil {
		log.Fatal("Cookie not found")
	} else {
		fmt.Fprintln(w, c2)
	}

}

func abundance(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "old-cookie",
		Value: "old-value",
	})
}

// We can set multiple cookies for each state
