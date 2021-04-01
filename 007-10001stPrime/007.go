package euler007

import (
	"math"
	"strconv"
)

// PrimeByPos returns the nth largest prime number
func PrimeByPos(pos int) int {

	// Find the max number that pos number can possibly be
	var max int
	if pos < 6 {
		max = 13
	} else {
		f := float64(pos)
		f = f * (math.Log(f) + math.Log(math.Log(f)))
		max = int(math.Ceil(f))
	}

	// Following the algorithm for the Sieve of Eratosthenes, we create a slice
	// of booleans that represents integers from 0 to max. This will be used to
	// map out if they are prime (false) or not prime (true).
	s := make([]bool, max)

	// 0 and 1 are not prime
	s[0] = true
	s[1] = true

	// Find all values that are the product of two numbers (i and j) less than
	// max. These values are not prime.
	for i := 2; i <= max; i++ {
		for j := 2; i*j < max; j++ {
			s[i*j] = true
		}
	}

	// All untouched booleans are prime. So we iterate over the array, making
	// note of the position where the prime number was found.
	primes := []int{2}
	for id, val := range s {
		if !val {
			primes = append(primes, id)
		}
	}

	return primes[pos]
}

func Solve() string {
	return strconv.Itoa(PrimeByPos(10001))
}
