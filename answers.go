package main

import (
	"fmt"

	euler1 "github.com/jeffre/euler/001-SumMultiples"
	euler2 "github.com/jeffre/euler/002-SumEvenFibonacci"
	euler3 "github.com/jeffre/euler/003-LargePrimeFactor"
	euler4 "github.com/jeffre/euler/004-LargestPalindromeProduct"
)

func main() {
	// Problem 1
	answer1 := euler1.MultipleSum{
		Maximum:  1000,
		Divisors: []int{3, 5},
	}.Calc()
	fmt.Printf("Problem 1: %v\n", answer1) //233168

	// Problem 2
	answer2 := euler2.SumEvenFibronacci(4000000)
	fmt.Printf("Problem 2: %v\n", answer2) //4613732

	// Problem 3
	answer3 := euler3.LargestPrimeFactor(600851475143)
	fmt.Printf("Problem 3: %v\n", answer3) //6857

	// Problem 2
	answer4 := euler4.LargestPalindromeProduct(999).Palindrome
	fmt.Printf("Problem 4: %v\n", answer4) //906609
}
