package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type country struct {
	Name string
	Capital string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	
	bd := country{
		Name: "Bangladesh",
		Capital: "Dhaka",
	}
	err := tpl.Execute(os.Stdout, bd)
	if err != nil {
		log.Fatalln(err)
	}
}