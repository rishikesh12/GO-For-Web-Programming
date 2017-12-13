package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	var s string
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {

		//open the file

		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		fmt.Println("file:", f, "error:", err, "header:", h) //printing the header information
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		s = string(bs)
	}
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `
      <form method="POST" enctype="multipart/form-data">
      <input type="file" name="q">
      <input type="submit">
      </form> </br>`+s)

}
