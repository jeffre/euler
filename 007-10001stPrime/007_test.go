package euler007

import "testing"

func TestPrimeByPos(t *testing.T) {
	cases := map[int]int{
		1:     2,
		2:     3,
		3:     5,
		4:     7,
		5:     11,
		6:     13,
		10001: 104743,
	}

	for pos, want := range cases {
		got := PrimeByPos(pos)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
