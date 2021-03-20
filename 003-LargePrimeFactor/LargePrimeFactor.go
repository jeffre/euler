package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

type Primes []int

var ErrNoPrimes = errors.New("no prime factors found")

func (p Primes) Last() (int, error) {
	if len(p) == 0 {
		return 0, ErrNoPrimes
	}
	return p[len(p)-1], nil
}

func PrimeFactors(n int) (primes Primes) {
	maximum := math.Ceil(math.Sqrt((float64(n))))

	for i := 2; i < int(maximum); i++ {
		for n%i == 0 {
			primes = append(primes, i)
			n = n / i
		}
	}

	return primes
}

func main() {

	n := 600851475143
	primes := PrimeFactors(n)

	largest_prime, err := primes.Last()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Largest prime factor of %v is %v", n, largest_prime)
}
