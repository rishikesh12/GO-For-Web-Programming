package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var fm = template.FuncMap{"fdate": monthDayYear}

func monthDayYear(t time.Time) string {
	return t.Format("02 - 01 - 2006") //We can also use constants in time pkg (ex-Kitchen)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("01_Templates/06_FunctionsInTemplates/02_BasicExample/tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
