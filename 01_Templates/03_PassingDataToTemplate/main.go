package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("01_Templates/03_PassingDataToTemplate/tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Krishna Ravishankar")
	if err != nil {
		log.Fatalln(err)
	}
}
