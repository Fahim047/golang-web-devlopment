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
	// capital := map[string]string{"Bangladesh": "Dhaka", "India": "New Delhi", "Pakistan": "Islamabad", "Nepal": "Kathmandu",}

	capital := map[string]string{
		"Bangladesh": "Dhaka", 
		"India": "New Delhi", 
		"Pakistan": "Islamabad", 
		"Nepal": "Kathmandu",
	}
	err := tpl.Execute(os.Stdout, capital)
	if err != nil {
		log.Fatalln(err)
	}
}