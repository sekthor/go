package main

import (
	"embed"
	"html/template"
	"log"
	"os"
)

var (
	//go:embed template.html
	tmplFS embed.FS
)

type TmplData struct {
	Username string
}

func main() {
	tmpls, err := template.ParseFS(tmplFS, "*.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl := tmpls.Lookup("template.html")

	if tmpl == nil {
		log.Fatal("could not find template")
	}

	data := TmplData{Username: "sekthor"}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal("could not execute template")
	}
}
