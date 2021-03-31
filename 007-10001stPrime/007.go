package euler007

import "strconv"

// PrimeByPos returns the nth largest prime number
func PrimeByPos(pos int) int {

	// Create a slice of booleans that represents integers from 0 to max. This
	// will be used to map out if they are prime (false) or not prime (true).
	max := 999999
	p := make([]bool, max)

	// 0 and 1 are not prime
	p[0] = true
	p[1] = true

	// Find all values that are the product of two numbers (i and j) less than
	// max. These values are not prime.
	for i := 2; i <= max; i++ {
		for j := 2; i*j < max; j++ {
			p[i*j] = true
		}
	}

	// All untouched booleans are prime. So we iterate over the array, making
	// note of the position where the prime number was found.
	primes := []int{2}
	for id, val := range p {
		if !val {
			primes = append(primes, id)
		}
	}

	return primes[pos]
}

func Solve() string {
	return strconv.Itoa(PrimeByPos(10001))
}
