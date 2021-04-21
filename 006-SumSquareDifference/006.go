package euler6

// SumOfSquares squares all numbers less than max and then returns their sum.
func SumOfSquares(max int) (total int) {
	for i := 1; i <= max; i++ {
		total += i * i
	}
	return total
}

// SquareOfSums sums all numbers less than max and then returns their square.
func SquareOfSums(max int) (total int) {
	for i := 1; i <= max; i++ {
		total += i
	}
	return total * total
}

// SumSquareDiff returns the difference of SquareOfSums and SumOfSquares using
// max.
func SumSquareDiff(max int) (total int) {
	sum := SumOfSquares(max)
	square := SquareOfSums(max)
	return square - sum
}

// Solve returns the answer for this Project Euler challenge
func Solve() int {
	return SumSquareDiff(100)
}
