package euler001

// SumMutiples sums all divisors less than max
func SumMultiples(max int, divisors []int) (total int) {
	for i := 1; i < max; i++ {
		for _, d := range divisors {
			if i%d == 0 {
				total += i
				break
			}
		}
	}
	return total
}
