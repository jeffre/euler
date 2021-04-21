package euler11

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGridVariable(t *testing.T) {
	cases := []struct {
		row  int
		col  int
		want int
	}{
		{0, 0, 8},
		{1, 1, 49},
		{19, 19, 48},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("row %v, col %v", test.row, test.col), func(t *testing.T) {
			got := grid[test.row][test.col]
			if got != test.want {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

func TestRight(t *testing.T) {
	cases := []struct {
		row  int
		col  int
		want []int
	}{
		{0, 0, []int{8, 2, 22, 97}},
		{1, 1, []int{49, 99, 40, 17}},
		{19, 19, []int{48, 0, 0, 0}},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("row %v, col %v", test.row, test.col), func(t *testing.T) {
			got := Right(test.row, test.col)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}
func TestDown(t *testing.T) {
	cases := []struct {
		row  int
		col  int
		want []int
	}{
		{0, 0, []int{8, 49, 81, 52}},
		{1, 1, []int{49, 49, 70, 31}},
		{19, 19, []int{48, 0, 0, 0}},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("row %v, col %v", test.row, test.col), func(t *testing.T) {
			got := Down(test.row, test.col)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

func TestDownRight(t *testing.T) {
	cases := []struct {
		row  int
		col  int
		want []int
	}{
		{0, 0, []int{8, 49, 31, 23}},
		{1, 1, []int{49, 31, 23, 51}},
		{19, 19, []int{48, 0, 0, 0}},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("row %v, col %v", test.row, test.col), func(t *testing.T) {
			got := DownRight(test.row, test.col)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

func TestDownLeft(t *testing.T) {
	cases := []struct {
		row  int
		col  int
		want []int
	}{
		{0, 0, []int{8, 0, 0, 0}},
		{1, 1, []int{49, 81, 0, 0}},
		{19, 19, []int{48, 0, 0, 0}},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("row %v, col %v", test.row, test.col), func(t *testing.T) {
			got := DownLeft(test.row, test.col)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	got := Solve()
	want := 70600674
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
