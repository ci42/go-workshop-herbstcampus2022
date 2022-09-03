package concurrency

import (
	"io"
	"os"
	"testing"
	"time"
)

// just a test that is formatted incorrectly
func TestBlockingChannel(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	c := make(chan string)

	blockingChannel(c)

	<-time.After(1 * time.Second)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	expected := "foobar\n"
	if string(out) != expected {
		t.Errorf("Fehler: %q != %q", out, expected)
	}
}
