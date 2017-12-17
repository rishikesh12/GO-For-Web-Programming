package main

import (
	"net/http"

	"github.com/satori/go.uuid"
)

//function to return the user
func getuser(w http.ResponseWriter, r *http.Request) user {
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}

	}
	http.SetCookie(w, c)
	var u user
	//check whether the user exists in the database
	if un, ok := dbsession[c.Value]; ok {
		u = dbuser[un]
	}

	return u

}

//function to check whether the user is loggedin
func alreadyregistered(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := dbsession[c.Value]
	_, ok := dbuser[un]
	return ok
}
