package euler003

import (
	"math"
)

// PrimeFactors returns the prime factors of n in ascending order. If no prime
// factors are found then nil is returned instead.
func PrimeFactors(n int) (primes []int) {

	// There are no integers above the square root of n that can be multiplied
	// together to get n.
	maximum := int(math.Sqrt((float64(n))))

	for i := 2; i <= maximum; i++ {

		// Because the same prime can reoccur, keep testing until we have
		// them all
		for n%i == 0 {
			primes = append(primes, i)
			n /= i
		}
	}

	return primes
}

// LargestPrimeFactor returns the single largest prime factor of n. If n is
// prime then it returns 0
func LargestPrimeFactor(n int) int {

	pf := PrimeFactors(n)

	if pf == nil {
		return 0
	}

	// The largest item in the slice is also the last one
	return pf[len(pf)-1]
}
