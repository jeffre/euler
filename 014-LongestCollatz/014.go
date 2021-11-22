package euler14

import (
	"math"
)

// BruteLongestCollatz finds the number that results in the longest collatz
// sequence.
func BruteforceLongestCollatz(max int) int {

	answer := 0
	longestChain := 0

	for i := 1; i <= max; i++ {
		num := i
		length := 1

		// Loop until collatz sequence completes
		for num != 1 {
			if num%2 == 0 {
				num /= 2
			} else {
				num = 3*num + 1
			}

			length++
		}

		// Check if this is the leader
		if length > longestChain {
			longestChain = length
			answer = i
		}
	}

	return answer
}

// cache stores the collatz length for a given number
var cache = map[int]int{1: 1}

// collatzLen returns the length for a given number's collatz sequence and
// contributes to cache
func collatzLen(n int) int {
	if v, ok := cache[n]; ok {
		return v
	}

	if n%2 == 0 {
		cache[n] = 1 + collatzLen(n/2)
	} else {
		// If n is odd, then 3n + 1 will be even. Therefore we can safely
		// combine two math steps into one
		cache[n] = 2 + collatzLen((3*n+1)/2)
	}

	return cache[n]
}

// LongestCollatz returns the number that produces the longest collatz sequence.
// This function is ~6 times faster than bruteforcing.
func LongestCollatz(max int) int {
	longestChain := 0
	answer := -1

	// Exclude the bottom half of possible solutions because the the n/2 rules
	// gives the top half an extra step
	min := int(math.Max(float64(max)/2, 1.0))

	for i := min; i <= max; i++ {
		num := i
		length := collatzLen(num)

		if length > longestChain {
			longestChain = length
			answer = num
		}
	}

	return answer
}

func Solve() int {
	return LongestCollatz(999999)
}
