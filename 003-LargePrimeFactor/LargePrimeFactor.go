package euler003

import (
	"math"
)

func PrimeFactors(n int) (primes []int) {

	// There are no integers above the square root of n that can be multiplied
	// together to get n.
	maximum := int(math.Sqrt((float64(n))))

	for i := 2; i <= maximum; i++ {

		// If all prime factors have been found than n will have been divided down to 0
		if n == 0 {
			break
		}

		// Test if i is a prime factor
		if n%i == 0 {
			primes = append(primes, i)
			n /= i

			// Deincrement the loop because n may have i more than once
			i--
		}
	}

	return primes
}
