package euler007

import (
	"math"
)

// Primes returns an ordered slice of primes up to size max
func Primes(max int) (primes []int) {

	// Following the algorithm for the Sieve of Eratosthenes, we create a slice
	// of booleans that represents integers from 0 to max. This will be used to
	// map out if they are prime (false) or not prime (true).
	s := make([]bool, max)

	// Find all values that are the product of two numbers (i and j) greater
	// than 1 and less than max. These values are not prime.
	for i := 0; i < max; i++ {
		if i <= 1 {
			s[i] = true
			continue
		}
		for j := 2; i*j < max; j++ {
			s[i*j] = true
		}
	}

	// All untouched booleans are prime. So we iterate over the array, making
	// note of the position where the prime number was found.
	for id, val := range s {
		if !val {
			primes = append(primes, id)
		}
	}

	return primes
}

// PrimeByPos returns the nth largest prime number
func PrimeByPos(n int) int {

	// Find the maximum size that position n can possibly be
	var max int
	if n < 6 {
		max = 13
	} else {
		f := float64(n)
		f = f * (math.Log(f) + math.Log(math.Log(f)))
		max = int(math.Ceil(f))
	}

	primes := Primes(max)
	return primes[n-1]
}

// Solve returns the answer for this Project Euler challenge
func Solve() int {
	return PrimeByPos(10001)
}
