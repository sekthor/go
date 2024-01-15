package main

import (
	"os"
	"text/template"
	"time"
)

type Person struct {
	Firstname string
	Lastname  string
	DoB       time.Time
}

func main() {

	person := Person{
		Firstname: "John",
		Lastname:  "Doe",
		DoB:       time.Now(),
	}

	path := "person.tmpl"
	tmpl, err := template.New(path).ParseFiles(path)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, person)
	if err != nil {
		panic(err)
	}
}
