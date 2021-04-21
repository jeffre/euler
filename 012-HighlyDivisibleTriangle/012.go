package euler12

import euler7 "github.com/jeffre/euler/007-10001stPrime"

// DivisorsCnt return the number of divisors for a given number n, provided a
// slice of prime numbers. Refer to Euler7 for a prime number generator.
func DivisorsCnt(n int, primes []int) (divisorcnt int) {

	/*
		Any integer N can be expressed as follows:
			N = P₁^A₁ * P₂^A₂ * P₃^A₃ * …
		where Pₙ is a prime number and Aₙ is its exponent.

		For example 28 = 2² * 7¹

		Furthermore, the number of divisors D(N) of any integer N can be computed from:
		    D(N) = (A₁+1) * (A₂+1) * (A₃+1) * …
		where Aₙ are the exponents of the distinct prime numbers which are factors of N

		For example, the number of divisors of 28 would thus be:
		    D(28) = (2+1) * (1+1) = 3 * 2 = 6
	*/

	// Starts at 1 because the calculation method uses multiplication
	divisorcnt = 1

	// Look for prime factors of `n`
	for _, p := range primes {

		// pfexponent (prime factor exponent) holds the exponential value of
		// a distinct prime factor of `n`.
		pfcnt := 0
		for n%p == 0 {
			n /= p
			pfcnt += 1
		}

		// Multiply this prime factor (plus one) with any previous distinct
		// prime factors
		if pfcnt > 0 {
			divisorcnt *= (pfcnt + 1)
		}

		// Test if all prime factors have been divided away
		if n == 1 {
			break
		}
	}
	return
}

// HighlyDivisibleTriangle returns the first triangle number to have `n`
// divisors.
func HighlyDivisibleTriangle(n int) (triangle int) {

	// Quickly handle these tricky divisors
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}

	// Generate a list of all primes less than 65500
	primes := euler7.Primes(65500)

	// Iterate through triangle numbers starting from 1
	for i := 1; ; i++ {
		triangle += i
		if DivisorsCnt(triangle, primes) >= n {
			return
		}
	}
}

// Solve returns the answer for this Project Euler challenge
func Solve() int {
	return HighlyDivisibleTriangle(500)
}
