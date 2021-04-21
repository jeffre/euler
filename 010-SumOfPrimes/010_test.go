package euler10

import (
	"fmt"
	"testing"
)

func TestSumPrimes(t *testing.T) {
	cases := map[int]int{
		0:       0,
		10:      17,
		2000000: 142913828922,
	}

	for max, want := range cases {
		t.Run(fmt.Sprintf("max %v", max), func(t *testing.T) {
			got := SumPrimes(max)
			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func BenchmarkSumPrimes(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		SumPrimes(2000000)
	}
}
