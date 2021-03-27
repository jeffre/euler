package euler003

import (
	"math"
)

func PrimeFactors(n int) (primes []int) {

	// There are no integers above the square root of n that can be multiplied
	// together to get n.
	maximum := int(math.Sqrt((float64(n))))

	for i := 2; i <= maximum; i++ {

		// Because the same integer can occur more than once as a prime factor
		// we will repeat the test while deincrementing n
		for n%i == 0 {
			primes = append(primes, i)
			n /= i
		}
	}

	return primes
}
