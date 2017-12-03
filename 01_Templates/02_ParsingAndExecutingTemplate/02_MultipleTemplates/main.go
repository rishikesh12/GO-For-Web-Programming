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

	tpl1, err1 := template.ParseGlob("01_Templates/02_ParsingAndExecutingTemplate/Text/*") //Glob is used to parse all the content within a folder
	if err1 != nil {
		log.Fatalln(err1)
	}

	err1 = tpl1.Execute(os.Stdout, nil)
	if err1 != nil {
		log.Fatalln(err1)
	}

	err1 = tpl1.ExecuteTemplate(os.Stdout, "sample2.txt", nil)
	if err1 != nil {
		log.Fatalln(err1)
	}
}

// we can use template.Must to perform error checking

/* func init(){
tpl:=template.Must(template.ParseGlob("Folder"))
}

This makes your code more performant
and we can declare tpl at package level scope

var tpl *template.Template
*/
