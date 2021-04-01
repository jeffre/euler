package euler006

func SumOfSquares(max int) (total int) {
	for i := 1; i <= max; i++ {
		total += i * i
	}
	return total
}

func SquareOfSums(max int) (total int) {
	for i := 1; i <= max; i++ {
		total += i
	}
	return total * total
}

func SumSquareDiff(max int) (total int) {
	sum := SumOfSquares(max)
	square := SquareOfSums(max)
	return square - sum
}

func Solve() int {
	return SumSquareDiff(100)
}
