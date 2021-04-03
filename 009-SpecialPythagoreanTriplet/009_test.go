package euler009

import (
	"fmt"
	"testing"
)

func TestPythagTriplet(t *testing.T) {
	cases := map[int]int{
		12:   60,
		1000: 31875000,
	}
	for sum, prod := range cases {
		t.Run(fmt.Sprintf("sum of %v", sum), func(t *testing.T) {
			got, _ := PythagTriplet(sum)
			if got != prod {
				t.Errorf("got %v want %v", got, prod)
			}
		})
	}

	t.Run("no match error", func(t *testing.T) {
		_, err := PythagTriplet(3)
		if err != ErrNoMatch {
			t.Errorf("got %q want %q", err, ErrNoMatch)
		}
	})
}

func BenchmarkPythagTriplet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PythagTriplet(1000)
	}
}
