package euler009

import (
	"testing"
)

func TestPythagTriplet(t *testing.T) {
	cases := map[int]int{
		12:   60,
		1000: 31875000,
	}
	for sum, prod := range cases {
		got := PythagTriplet(sum)
		if got != prod {
			t.Errorf("got %v want %v", got, prod)
		}
	}
}

func BenchmarkPythagTriplet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PythagTriplet(1000)
	}
}
