package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("01_Templates/04_VariablesInTemplate/tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Welcome to my world")
	if err != nil {
		log.Fatalln(err)
	}
}
