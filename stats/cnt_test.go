package stats

import (
	"testing"
)

func TestImageType(t *testing.T) {
	c := Cnt{Source: "Pixabay", Total: 28}

	c.Source = "Pexels" // dummy operation to avoid compile error of unused var
}
