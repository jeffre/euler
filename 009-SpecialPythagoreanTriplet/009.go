package euler9

import (
	"errors"
	"math"
)

var ErrNoMatch = errors.New("pythagorean triplet not found")

// PythagTriplet takes the sum of the 3 unique integers of a pythagorean
// triplet and returns a product of a triplet that has the same sum
func PythagTriplet(s int) (int, error) {

	// Iterate backward from the maximum possible value for c which is s minus
	// 7 (7 being the sum of the two smallest members of the smallest possible
	// pythagorean triple 3² + 4² = 5²)
	for i := s - 7; i > 0; i-- {
		c := i

		// Iterate backward from the remainder of the sum - c - 1. The minus 1
		// prevents 'a' from being set to 0. The conditional '(s-c)/2' ensures
		// that b will always be large than a.
		for j := s - c - 1; j > (s-c)/2; j-- {
			b := j

			// 'a' then becomes the smallest triplet
			a := s - c - b

			// Check if a² + b² = c²
			if math.Pow(float64(a), 2)+math.Pow(float64(b), 2) == math.Pow(float64(c), 2) {
				return int(a * b * c), nil
			}
		}
	}

	// Failed to get an answer
	return 0, ErrNoMatch
}

// Solve returns the answer for this Project Euler challenge
func Solve() int {
	a, _ := PythagTriplet(1000)
	return a
}

/*
This section contains my scratch notes which might be helpful later on in
finding an algorithm.

Given:
    a+b+c=12      a²+b²=c²
	c=12-a-b      (a²+b²)/c=c
	    12-a-b = (a²+b²)/c
	c(12-a-b) = (a²+b²)
	c = (a²+b²)/(12-a-b)

Testing with a=3,b=4,c=5
	c = (9+16)/(12-3-4) = 25/5 = 5
*/
