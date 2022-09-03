package helloworld

import (
	"io"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	helloWorld()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	expected := "Hello Herbstcampus 2022!\n"
	if string(out) != expected {
		t.Errorf("Fehler: %q != %q", out, expected)
	}
}
