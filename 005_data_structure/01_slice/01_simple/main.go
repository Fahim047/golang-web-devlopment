package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	name := []string{"Fahim", "Rian", "Abu Bokor", "Rabbe"}
	err := tpl.Execute(os.Stdout, name)
	if err != nil {
		log.Fatalln(err)
	}
}