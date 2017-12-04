package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("01_Templates/05_PassingCompositeDataIntoTemplates/02_PassingMapContent/tpl.gohtml"))
}

func main() {
	dost := map[string]string{
		"India":  "Krishna",
		"Canada": "Charan",
		"USA":    "Chethan",
	}

	err := tpl.Execute(os.Stdout, dost)
	if err != nil {
		log.Fatalln(err)
	}
}
