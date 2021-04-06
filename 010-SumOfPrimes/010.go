package euler010

// SumPrimes returns the sum of all primes less than max
func SumPrimes(max int) (sum int) {

	// Following the algorithm for the Sieve of Eratosthenes, we create a slice
	// of booleans that represents integers from 0 to max. This will be used to
	// map out if they are prime (false) or not prime (true).
	s := make([]bool, max)

	// Find all values that are the product of two numbers (i and j) greater
	// than 1 and less than max. These values are not prime.
	for i := 1; i < max; i++ {
		if i == 1 {
			s[i] = true
			continue
		}
		for j := 2; i*j < max; j++ {
			s[i*j] = true
		}
	}

	// All untouched (false) booleans are prime. So we iterate over the array,
	// summing the position where the prime number was found.
	for id, val := range s {
		if !val {
			sum += id
		}
	}

	return sum
}

func Solve() int {
	return SumPrimes(2000000)
}
