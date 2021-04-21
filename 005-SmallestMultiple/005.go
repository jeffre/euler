package euler5

import (
	"math"

	euler7 "github.com/jeffre/euler/007-10001stPrime"
)

// SmallestMultiple finds the smallest common multiple of integers that are less
// than max.
func SmallestMultiple(max int) (sm int) {

	// Smallest Multiple will be multiplied against other numbers so set
	// default value to 1.
	sm = 1

	// Get all primes numbers that do not exceed max
	primes := euler7.Primes(max)

	// Iterate over the primes
	for _, prime := range primes {

		// Calculate the number of times this prime will occur as a factor within the
		// sequence of 2 thru `max`.
		exponent := int(math.Log(float64(max)) / math.Log(float64(prime)))

		// Calculate prime^exponent and multiply this with the prior prime^exponents
		sm *= int(math.Pow(float64(prime), float64(exponent)))
	}

	return sm
}

// Solve returns the answer for this Project Euler challenge
func Solve() int {
	return SmallestMultiple(20)
}

/*
// Left for reference only. Using bruteforcing to solve this problem is slower
// (by magnitudes) than the current method.

// cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
// BenchmarkSmallestMultipleBrute-8               5         202224960 ns/op
// BenchmarkSmallestMultiple-8              2173063               544.4 ns/op

// SmallestMultiple finds the smallest common multiple of integers that are less
// than max.
func SmallestMultipleBrute(max int) int {

	// Iterate into infinity in increments of the largest divisor
	for i := max; ; i += max {

		// Presume i is a multiple until proven otherwise in a modulus test
		isMultiple := true
		for j := max; j > 0; j-- {
			if i%j != 0 {
				isMultiple = false
				break
			}
		}

		// At this point either all divisors are multiples or one of them is
		// not and has flipped the boolean back to false.
		if isMultiple {
			return i
		}
	}
}
*/
