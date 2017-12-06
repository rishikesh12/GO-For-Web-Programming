package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("01_Templates/07_PredefinedGlobalFunctions/01_BasicExample/tpl.gohtml"))
}
func main() {
	s := []string{"one", "two", "three", "four"}
	err := tpl.Execute(os.Stdout, s)
	if err != nil {
		log.Fatalln(err)
	}
}
