package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session-id")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(),
			//Secure:true, Https option
			HttpOnly: true, //Can't access the cookie with javascript
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}

//We are using satori go.uuid package (which is a third party package)
