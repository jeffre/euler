package euler003

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrimeFactors(t *testing.T) {

	tests := map[int][]int{
		8:            {2, 2, 2},
		17:           nil,
		13195:        {5, 7, 13, 29},
		600851475143: {71, 839, 1471, 6857},
	}

	for n, want := range tests {
		t.Run(fmt.Sprintf("Testing %v", n), func(t *testing.T) {
			got := PrimeFactors(n)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v want %+v", got, want)
			}
		})
	}
}

func TestLargestPrimeFactor(t *testing.T) {

	tests := map[int]int{
		0:            0,
		1:            0,
		2:            0,
		3:            0,
		4:            2,
		8:            2,
		17:           0,
		13195:        29,
		600851475143: 6857,
	}

	for n, want := range tests {
		t.Run(fmt.Sprintf("Testing %v", n), func(t *testing.T) {
			got := LargestPrimeFactor(n)
			if got != want {
				t.Errorf("got %+v want %+v", got, want)
			}
		})
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeFactors(600851475143)
	}
}
