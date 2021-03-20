package main

import "fmt"

func SumEvenFibronacci(max int) (total int) {

	// Early return for odd these odd cases
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

func main() {
	max := 4000000
	fmt.Println(SumEvenFibronacci(max))
}
