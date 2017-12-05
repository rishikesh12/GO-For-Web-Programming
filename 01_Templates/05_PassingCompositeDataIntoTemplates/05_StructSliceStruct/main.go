package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Company string
	Model   string
}

type item struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("01_Templates/05_PassingCompositeDataIntoTemplates/05_StructSliceStruct/tpl.gohtml"))
}

func main() {
	buddha := sage{Name: "Buddha", Motto: "Life is a mystery"}
	bose := sage{Name: "Bose", Motto: "Give me blood I shall give you freedom"}

	sages := []sage{buddha, bose}

	i10 := car{Company: "Hyndai", Model: "sportz"}
	breeza := car{Company: "Maruthi", Model: "spectrum"}

	cars := []car{i10, breeza}

	newitem := item{Wisdom: sages, Transport: cars}
	err := tpl.Execute(os.Stdout, newitem)
	if err != nil {
		log.Fatalln(err)
	}
}
