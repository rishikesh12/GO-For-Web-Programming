package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("01_Templates/05_PassingCompositeDataIntoTemplates/tpl.gohtml"))
}

func main() {
	friends := []string{"Krishna", "Charan", "Chethan", "Harsh"}
	err := tpl.Execute(os.Stdout, friends) //(writer,data)
	if err != nil {
		log.Fatalln(err)
	}

}

//Find out how to create a new html?
