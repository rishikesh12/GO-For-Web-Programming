package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("01_Templates/02_ParsingAndExecutingTemplate/krishkrash.gohtml") //template is a container holding the content of the parsed files
	if err != nil {
		log.Fatalln(err)
	}
	nf, err := os.Create("01_Templates/02_ParsingAndExecutingTemplate/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(nf, nil)
}

//ParseFiles returns a pointer of template type
//Execute takes a writer interface and since a file implements writer interface we can pass the file as a writer
