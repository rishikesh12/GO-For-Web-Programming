package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", count)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func count(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}
	count, err := strconv.Atoi(cookie.Value) //ascii to integer
	if err != nil {
		log.Fatal(err)
	}
	count++
	cookie.Value = strconv.Itoa(count) //integer to ascii

	http.SetCookie(w, cookie)

	fmt.Fprintln(w, "Count Value", cookie.Value)

}

//program to count number of visits by a user
