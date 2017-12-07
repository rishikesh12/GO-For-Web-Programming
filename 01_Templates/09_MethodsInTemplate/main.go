package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func (p person) Double() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return p.Age + x
}

func init() {
	tpl = template.Must(template.ParseFiles("01_Templates/09_MethodsInTemplate/tpl.gohtml"))
}
func main() {
	p1 := person{Name: "Krishna", Age: 21}
	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
