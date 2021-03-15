package main

import (
	"fmt"
)

// Divisors are used to test for mutliplicity
type Divisors []int

// SumMultiples sums all integers below `max` and a multiple of a `divisors`
func SumMultiples(max int, divisors Divisors) (total int) {

	for i := 1; i < max; i++ {
		for _, d := range divisors {
			if i%d == 0 {
				total += i
			}
		}
	}
	return total
}

func main() {
	msum := SumMultiples(1000, Divisors{3, 5})
	fmt.Println(msum)
}
