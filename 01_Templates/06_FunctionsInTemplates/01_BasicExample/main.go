package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

//FuncMap to register functions
//uc is toupper function from package strings
//ft is a function declared which returns the first three characters

var fm = template.FuncMap{"uc": strings.ToUpper, "ft": firstThree}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("01_Templates/06_FunctionsInTemplates/01_BasicExample/tpl.gohtml")) //Funcs call should be done before parsing
}
func main() {
	buddha := sage{Name: "Buddha", Motto: "Life is a mystery"}
	bose := sage{Name: "Bose", Motto: "Give me blood and I shall give you freedom"}
	sages := []sage{buddha, bose}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages) //have to use executetemplate use of Execute will cause an error
	if err != nil {
		log.Fatalln(err)

	}

}
