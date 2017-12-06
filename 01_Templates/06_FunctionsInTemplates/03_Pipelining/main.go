package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

var tpl *template.Template
var fp = template.FuncMap{"sq": square, "thr": three, "sqrt": squareroot}

func init() {
	tpl = template.Must(template.New("").Funcs(fp).ParseFiles("01_Templates/06_FunctionsInTemplates/03_Pipelining/tpl.gohtml"))
}

func square(n int) int {
	return n * n
}

func three(n int) int {
	return n * n * n
}

func squareroot(n int) int {
	return int(math.Sqrt(float64(n)))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
