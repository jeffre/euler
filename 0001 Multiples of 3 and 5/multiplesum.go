package main

import (
	"fmt"
	"os"
	"strconv"
)

type divisors []int

// Multiples returns a slice of all integers below `maximum` that are multiples of `divisors`
func Multiples(maximum int, divisors divisors) []int {
	multiples := []int{}
	for i := 1; i < maximum; i++ {
		for _, d := range divisors {
			if i%d == 0 {
				multiples = append(multiples, i)
				break
			}
		}
	}
	return multiples
}

// Sum sums all integers contained in a slice
func Sum(integers []int) int {
	total := 0
	for _, i := range integers {
		total += i
	}
	return total
}

// SumMultiples sums all integers below `maximum` that are multiples of `divisors`
func SumMultiples(maximum int, divisors divisors) int {
	multiples := Multiples(maximum, divisors)
	return Sum(multiples)
}

func usage() {
	fmt.Printf("Usage: %s MAX DIVISOR [DIVISOR [...]]\n", os.Args[0])
}

func main() {

	if len(os.Args) <= 2 {
		usage()
		os.Exit(1)
	}

	rawArgs := os.Args[1:]

	// Convert command line args (strings) to ints
	var err error
	intArgs := make([]int, len(rawArgs))
	for i := 0; i < len(rawArgs); i++ {
		if intArgs[i], err = strconv.Atoi(rawArgs[i]); err != nil {
			fmt.Printf("%v\n", err)
			panic(err)
		}
	}

	max, nums := intArgs[0], intArgs[1:]
	multiples := Multiples(max, nums)
	fmt.Printf("Multiples are: %v\n", multiples)
	sum := Sum(multiples)
	fmt.Printf("Sum of multiples: %v\n", sum)
}
