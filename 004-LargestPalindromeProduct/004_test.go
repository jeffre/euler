package euler004

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		number  int
		reverse int
	}{
		{0, 0},
		{1, 1},
		{12, 21},
		{44, 44},
		{12345, 54321},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.number), func(t *testing.T) {
			got := Reverse(test.number)
			if got != test.reverse {
				t.Errorf("got %v want %v", got, test.reverse)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		number     int
		palindrome bool
	}{
		{0, true},
		{1, true},
		{12, false},
		{121, true},
		{1221, true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Testing %v", test.number), func(t *testing.T) {
			got := IsPalindrome(test.number)
			if got != test.palindrome {
				t.Errorf("got %v want %v", got, test.palindrome)
			}
		})
	}
}

func TestLargestPalindromeProduct(t *testing.T) {
	tests := []struct {
		max        int
		palindrome int
	}{
		{0, 0},
		{9, 9},
		{99, 9009},
		{999, 906609},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Testing %v", test.max), func(t *testing.T) {
			got := LargestPalindromeProduct(test.max)
			if got.Palindrome != test.palindrome {
				t.Errorf("got %v want %v", got.Palindrome, test.palindrome)
			}
		})
	}
}

func BenchmarkLargestPalindromeProduct(b *testing.B) {
	max := 999
	for i := 0; i < b.N; i++ {
		LargestPalindromeProduct(max)
	}
}
