package euler15

import (
	"fmt"
	"testing"
)

func TestLatticePath(t *testing.T) {
	cases := []struct {
		size   dimensions
		routes int
	}{
		{size: dimensions{1, 1}, routes: 2},
		{size: dimensions{2, 2}, routes: 6},
		{size: dimensions{3, 3}, routes: 20},
		//{size: dimensions{4, 4}, routes: 17},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%+v", test.size), func(t *testing.T) {
			got := LatticePath(test.size)
			if got != test.routes {
				t.Errorf("got %v want %v", got, test.routes)
			}
		})
	}
}
