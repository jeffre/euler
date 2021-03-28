package euler001

import (
	"fmt"
	"testing"
)

func TestSumMultiples(t *testing.T) {
	cases := []struct {
		max      int
		divisors []int
		want     int
	}{
		{0, []int{3, 5}, 0},
		{10, []int{3, 5}, 23},
		{1000, []int{3, 5}, 233168},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%+v", test), func(t *testing.T) {
			got := SumMultiples(test.max, test.divisors)
			if got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}
