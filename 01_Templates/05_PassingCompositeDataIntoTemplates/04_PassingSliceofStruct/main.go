package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("01_Templates/05_PassingCompositeDataIntoTemplates/04_PassingSliceofStruct/tpl.gohtml"))
}

func main() {
	buddha := sage{Name: "Buddha", Motto: "Life is a mystery"}
	bose := sage{Name: "Bose", Motto: "Give me blood i will give you freedom"}

	sages := []sage{buddha, bose}

	err := tpl.Execute(os.Stderr, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
