package stats

import (
	"fmt"
)

func ExampleStatsInc() {
	s := Stats{}

	s.Inc("Pixabay")
	s.Inc("Pixabay")
	s.Inc("Pixabay")
	s.Inc("Unsplash")

	fmt.Printf("%s", s)

	// Output:
	// Pixabay: 3
	// Unsplash: 1

}
