package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles(`01_Templates/07_PredefinedGlobalFunctions/03_BasicExample(comparison)/tpl.gohtml`))
}

func main() {

	g1 := struct { //Alternate way to declare and initialize  a struct
		Score1 int
		Score2 int
	}{
		3,
		7,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
}
