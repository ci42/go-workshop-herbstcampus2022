package web

import (
	"fmt"
	"os"

	"text/template"
)

func NewTemplateFromFile(fileName string) (*template.Template, error) {
	src, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error parsing web/static/index.html: %s", err)
	}

	t, err := template.New("index.html").Parse(string(src))
	if err != nil {
		return nil, fmt.Errorf("error creating template for index.html: %s", err)
	}

	return t, nil
}
