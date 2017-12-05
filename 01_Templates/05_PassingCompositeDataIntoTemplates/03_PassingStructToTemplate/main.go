package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("01_Templates/05_PassingCompositeDataIntoTemplates/03_PassingStructToTemplate/tpl.gohtml"))
}

type sage struct {
	Name  string
	Motto string
}

func main() {
	budda := sage{Name: "Budda", Motto: "The belief of no beliefs"}
	err := tpl.Execute(os.Stdout, budda)
	if err != nil {
		log.Fatalln(err)
	}
}
