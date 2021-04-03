package euler009

import (
	"math"
)

// PythagTriplet takes the sum of the 3 unique integers of a pythagorean
// triplet and returns a product of a triplet that has the same sum
func PythagTriplet(s int) int {

	// Iterate backward from the maximum possible value for c which is s minus
	// 3 (this accomodates integers a and b plus the fact they they cannot have
	// the same value
	for i := s - 3; i > 0; i-- {
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
				return int(a * b * c)
			}
		}
	}

	// Failed to get an answer
	return 0
}

func Solve() int {
	return PythagTriplet(1000)
}

/*
This section contains my scratch notes which might be helpful later on in
finding an algorhythm.

Given:
    a+b+c=12      a²+b²=c²
	c=12-a-b      (a²+b²)/c=c
	    12-a-b = (a²+b²)/c
	c(12-a-b) = (a²+b²)
	c = (a²+b²)/(12-a-b)

Testing with a=3,b=4,c=5
	c = (9+16)/(12-3-4) = 25/5 = 5
*/
