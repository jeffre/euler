package euler005

import "strconv"

// SmallestMultiple find the smallest common multiple of integers that are less
// than max
func SmallestMultiple(max int) int {

	// Iterate into infinity in increments of the largest divisor
	for i := max; ; i += max {

		// Presume i is a multiple until proven otherwise in a modulus test
		isMultiple := true
		for j := max; j > 0; j-- {
			if i%j != 0 {
				isMultiple = false
				break
			}
		}

		// At this point either all divisors are multiples or one of them is
		// not and has flipped the boolean back to false.
		if isMultiple {
			return i
		}
	}
}

func Solve() string {
	max := 20
	answer := SmallestMultiple(max)
	return strconv.Itoa(answer)
}
