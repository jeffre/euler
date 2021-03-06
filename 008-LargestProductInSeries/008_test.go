package euler8

import "testing"

func TestLargestProductInSeries(t *testing.T) {
	cases := map[int]int{
		4:  5832,
		13: 23514624000,
	}
	for length, want := range cases {
		got := LargestProductInSeries(length)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func BenchmarkLargestProductInSeries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestProductInSeries(13)
	}
}

func TestLargestProductInSeries2(t *testing.T) {
	cases := map[int]int{
		4:  5832,
		13: 23514624000,
	}
	for length, want := range cases {
		got := LargestProductInSeries2(length)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func BenchmarkLargestProductInSeries2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestProductInSeries2(13)
	}
}
