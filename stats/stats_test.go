package stats

import "testing"

func TestStatsType(t *testing.T) {
	s := Stats{Cnt{"Pexels", 28}, Cnt{"Unsplash", 100}}

	s[0].Total = 0 // dummy operation to avoid compile error of unused var
}
