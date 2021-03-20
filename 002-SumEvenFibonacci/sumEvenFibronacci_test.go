package main

import (
	"fmt"
	"testing"
)

func TestSumEvenFibronacci(t *testing.T) {
	cases := map[int]int{
		0:       0,
		1:       0,
		2:       2,
		100:     44,
		4000000: 4613732,
	}
	for max, sumef := range cases {
		t.Run(fmt.Sprintf("maximum size %v", max), func(t *testing.T) {
			got := SumEvenFibronacci(max)

			if got != sumef {
				t.Errorf("got %v want %v", got, sumef)
			}
		})
	}
}
