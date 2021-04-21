package euler1

// SumMultiples sums all divisors less than max
func SumMultiples(max int, divisors []int) (total int) {

	// Check all integers less than max
	for i := 1; i < max; i++ {

		// Try each of the divisors. If any work, add it to the total and then
		// break out of the divisors loop
		for _, d := range divisors {

			if i%d == 0 {
				total += i
				break
			}
		}
	}
	return total
}

// Solve returns the answer for this Project Euler challenge
func Solve() int {
	max := 1000
	divisors := []int{3, 5}
	return SumMultiples(max, divisors)
}
