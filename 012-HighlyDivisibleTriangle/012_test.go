package euler012

import (
	"fmt"
	"testing"
)

func TestHighlyDivisibleTriangle(t *testing.T) {
	cases := []struct {
		divisorCount int
		triangleNum  int
	}{
		{0, 0},
		{1, 1},
		{2, 3},
		{4, 6},
		{5, 28},
		{6, 28},
		{500, 76576500},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("divisor %v", test.divisorCount), func(t *testing.T) {
			got := HighlyDivisibleTriangle(test.divisorCount)
			if got != test.triangleNum {
				t.Errorf("got %v want %v", got, test.triangleNum)
			}
		})
	}
}

func BenchmarkHighlyDivisibleTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HighlyDivisibleTriangle(501)
	}
}
