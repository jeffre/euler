package euler010

import (
	euler7 "github.com/jeffre/euler/007-10001stPrime"
)

// SumPrimes returns the sum of all primes less than max
func SumPrimes(max int) (sum int) {

	// Use prime number finder from euler 7
	primes := euler7.Primes(max)

	// Sum the primes
	for _, p := range primes {
		sum += p
	}

	return sum
}

func Solve() int {
	return SumPrimes(2000000)
}
