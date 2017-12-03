package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("01_Templates/02_ParsingAndExecutingTemplate/krishkrash.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("01_Templates/02_ParsingAndExecutingTemplate/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "krishkrash.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil) //will execute whichever template it finds first
	if err != nil {
		log.Fatalln(err)
	}
}
