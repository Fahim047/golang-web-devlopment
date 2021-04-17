package main

import (
	"log"
	"os"
	"text/template"
)


func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal("error while parsing file", err)
	}

	err = tpl.Execute(os.Stdout, err)
	if err != nil {
		log.Fatalln(err)
	}
}