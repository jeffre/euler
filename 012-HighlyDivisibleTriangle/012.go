package euler012

func HighlyDivisibleTriangle(divisors int) (triangle int) {

	// Quickly handle these tricky divisors
	switch divisors {
	case 0:
		return 0
	case 1:
		return 1
	}

	// Iterate through triangle numbers starting from 1
	for i := 1; ; i++ {
		triangle += i

		// Get the divisors of the triangle number
		var factors []int
		max := triangle
		for j := 1; j < max; j++ {

			// if j is a divisor
			if triangle%j == 0 {

				// Append first factor
				factors = append(factors, j)
				remainder := triangle / j

				// Reduce max
				max = remainder

				// Check if theres a unique second factor
				if remainder != j {
					factors = append(factors, max)
				}

				// Check if we've found desired triangle
				if len(factors) >= divisors {
					//fmt.Printf("%v: %+v\n", triangle, factors)
					return triangle
				}
			}
		}
	}
}

func Solve() int {
	return HighlyDivisibleTriangle(500)
}
