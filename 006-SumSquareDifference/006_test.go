package euler006

import "testing"

func TestSumOfSquares(t *testing.T) {
	cases := map[int]int{
		10:  385,
		100: 338350,
	}

	for max, want := range cases {
		got := SumOfSquares(max)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestSquareOfSums(t *testing.T) {
	cases := map[int]int{
		10:  3025,
		100: 25502500,
	}

	for max, want := range cases {
		got := SquareOfSums(max)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestSumSquareDiff(t *testing.T) {
	cases := map[int]int{
		10:  2640,
		100: 25164150,
	}

	for max, want := range cases {
		got := SumSquareDiff(max)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func BenchmarkSumSquareDiff(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		max := 100
		SumSquareDiff(max)
	}
}
