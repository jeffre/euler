package main

import (
	"reflect"
	"testing"
)

func TestMultiples(t *testing.T) {
	maximum := 10
	divisors := divisors{3, 5}

	got := Multiples(maximum, divisors)
	want := []int{3, 5, 6, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	got := Sum([]int{3, 5})
	want := 8
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumMultiples(t *testing.T) {
	got := SumMultiples(10, divisors{3, 5})
	want := 23
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
