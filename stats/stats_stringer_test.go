package stats

import "fmt"

func ExampleStats_String() {
	s := Stats{
		{"Pexels", 20},
		{"Unsplash", 100},
	}.String()

	fmt.Print(s)

	// Output:
	// Pexels: 20
	// Unsplash: 100
}
