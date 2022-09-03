package web

import (
	"log"
	"text/template"
)

var IndexTemplate *template.Template

func init() {
	var err error
	IndexTemplate, err = NewTemplateFromFile("./templates/index.html")
	if err != nil {
		log.Fatal("error creating template for index.html: ", err)
	}
}
