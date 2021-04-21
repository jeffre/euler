package euler13

import (
	"math/big"
	"testing"
)

func TestGetNums(t *testing.T) {
	got := getNums()[0]
	want, _ := new(big.Int).SetString("37107287533902102798797998220837590246510135740250", 10)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	nums := getNums()
	got := Sum(nums)
	want, _ := new(big.Int).SetString("5537376230390876637302048746832985971773659831892672", 10)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFirstTen(t *testing.T) {
	b, _ := new(big.Int).SetString("12345678901234567890", 10)
	got := FirstTen(b)
	want := 1234567890
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSolve(t *testing.T) {
	got := Solve()
	want := 5537376230
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
