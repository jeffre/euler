package euler5

import (
	"fmt"
	"testing"
)

func TestSmallestMultiple(t *testing.T) {
	cases := []struct {
		max  int
		want int
	}{
		{10, 2520},
		{20, 232792560},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("testing %v", test.max), func(t *testing.T) {
			got := SmallestMultiple(test.max)
			if got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})

	}
}

func BenchmarkSmallestMultiple(b *testing.B) {
	max := 20
	for i := 0; i < b.N; i++ {
		SmallestMultiple(max)
	}
}

/*
// This section is here for reference on how to test a solution by bruteforce.

func TestSmallestMultipleBrute(t *testing.T) {
	cases := []struct {
		max  int
		want int
	}{
		{10, 2520},
		{20, 232792560},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("testing %v", test.max), func(t *testing.T) {
			got := SmallestMultipleBrute(test.max)
			if got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})

	}
}

func BenchmarkSmallestMultipleBrute(b *testing.B) {
	max := 20
	for i := 0; i < b.N; i++ {
		SmallestMultipleBrute(max)
	}
}
*/
