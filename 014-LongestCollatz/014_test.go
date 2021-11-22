package euler14

import (
	"fmt"
	"testing"
)

var longestCases = []struct {
	max     int
	longest int
}{
	{1, 1},
	{2, 2},
	{13, 9},
	{999999, 837799},
}

func TestCollatzLen(t *testing.T) {
	cases := []struct {
		num    int
		length int
	}{
		{1, 1},
		{2, 2},
		{13, 10},
		{5000, 29},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%v", test.num), func(t *testing.T) {
			got := collatzLen(test.num)
			if got != test.length {
				t.Errorf("got %v want %v", got, test.length)
			}
		})
	}
}

func TestBruteforceLongestCollatz(t *testing.T) {
	for _, test := range longestCases {
		t.Run(fmt.Sprintf("%v", test.max), func(t *testing.T) {
			got := BruteforceLongestCollatz(test.max)
			if got != test.longest {
				t.Errorf("got %v want %v", got, test.longest)
			}
		})
	}
}

func TestLongestCollatz(t *testing.T) {
	for _, test := range longestCases {
		t.Run(fmt.Sprintf("%v", test.max), func(t *testing.T) {
			got := LongestCollatz(test.max)
			if got != test.longest {
				t.Errorf("got %v want %v", got, test.longest)
			}
		})
	}
}

var max = 99999

func BenchmarkBruteforceLongestCollatz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteforceLongestCollatz(max)
	}
}

func BenchmarkLongestCollatz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestCollatz(max)
	}
}
