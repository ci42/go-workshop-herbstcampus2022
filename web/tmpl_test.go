package web

import (
	"os"
	"testing"
)

func TestNewTemplateFromFile(t *testing.T) {
	_, err := NewTemplateFromFile("nonexistant")
	if err == nil {
		t.Errorf("NewTemplateFromFile schlägt nicht fehl, wenn es eine nicht existente Datei bekommt")
	}

	_, err = NewTemplateFromFile("web/tmpl_test.go")
	if err == nil {
		t.Errorf("NewTemplateFromFile schlägt nicht fehl, wenn ein defektes Template eingelesen wird")
	}

}

func ExampleNewTemplateFromFile() {
	t, _ := NewTemplateFromFile("./templates/test_template.txt")
	t.Execute(os.Stdout, struct{ Foo string }{"bar"})

	// Output: bar
}
