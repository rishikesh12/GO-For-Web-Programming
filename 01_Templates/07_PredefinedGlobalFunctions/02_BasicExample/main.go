package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("01_Templates/07_PredefinedGlobalFunctions/02_BasicExample/tpl.gohtml"))
}

type user struct {
	Name  string
	Admin bool
}

func main() {
	u1 := user{Name: "Krishna", Admin: false}
	u2 := user{Name: "Harsh", Admin: true}
	u3 := user{Name: "", Admin: false}
	users := []user{u1, u2, u3}
	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
