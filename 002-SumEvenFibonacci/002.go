package euler2

// SumEvenFibonacci returns the sum of even fibonacci numbers less than max
func SumEvenFibonacci(max int) (total int) {

	// Early return these odd cases
	switch max {
	case 0:
		return 0
	case 1:
		return 0
	}

	fib := [2]int{1, 1}

	for {
		next := fib[0] + fib[1]
		if next > max {
			break
		}
		if next%2 == 0 {
			total += next
		}
		fib[0] = fib[1]
		fib[1] = next
	}

	return total
}

// Solve returns the answer for this Project Euler challenge
func Solve() int {
	return SumEvenFibonacci(4000000)
}
