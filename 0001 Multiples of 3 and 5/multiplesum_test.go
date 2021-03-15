package main

import (
	"fmt"
	"testing"
)

func TestSumMultiples(t *testing.T) {
	cases := []struct {
		Max      int
		Divisors Divisors
		MSum     int
	}{
		{0, Divisors{}, 0},
		{10, Divisors{11}, 0},
		{10, Divisors{3, 5}, 23},
		{1000, Divisors{3, 5}, 266333},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			got := SumMultiples(test.Max, test.Divisors)
			if got != test.MSum {
				t.Errorf("got %v want %v", got, test.MSum)
			}

		})
	}
}
