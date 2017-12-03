package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Krishna"
	tpl := fmt.Sprintf(`<!DOCTYPE html>
        <html lang="en">
        <head>
        <title>Hello</title>
        </head>
        <body>
        <h1>` + name + `</h1>
        </body>
        </html>`)

	fp, err := os.Create("01_Templates/01_TemplatingWithConcatenation/index2.html")
	if err != nil {
		log.Fatalln("Error Creating File")
	}
	defer fp.Close()
	io.Copy(fp, strings.NewReader(tpl))
	fmt.Println(tpl)
}
