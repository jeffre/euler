package euler008

import (
	"math/big"
	"strconv"
)

const (
	NUM = "73167176531330624919225119674426574742355349194934" +
		"96983520312774506326239578318016984801869478851843" +
		"85861560789112949495459501737958331952853208805511" +
		"12540698747158523863050715693290963295227443043557" +
		"66896648950445244523161731856403098711121722383113" +
		"62229893423380308135336276614282806444486645238749" +
		"30358907296290491560440772390713810515859307960866" +
		"70172427121883998797908792274921901699720888093776" +
		"65727333001053367881220235421809751254540594752243" +
		"52584907711670556013604839586446706324415722155397" +
		"53697817977846174064955149290862569321978468622482" +
		"83972241375657056057490261407972968652414535100474" +
		"82166370484403199890008895243450658541227588666881" +
		"16427171479924442928230863465674813919123162824586" +
		"17866458359124566529476545682848912883142607690042" +
		"24219022671055626321111109370544217506941658960408" +
		"07198403850962455444362981230987879927244284909188" +
		"84580156166097919133875499200524063689912560717606" +
		"05886116467109405077541002256983155200055935729725" +
		"71636269561882670428252483600823257530420752963450"
)

// LargestProductInSeries returns the largest product of l adjacent digits
// This method iterates over NUM as a string
func LargestProductInSeries(l int) (m int) {

	digits := make([]int, l)
	max := 0

	for i := 0; i < len(NUM); i++ {

		// Add one digit to the tail of the slice
		d, err := strconv.Atoi(NUM[i : i+1])
		if err != nil {
			panic(err)
		}
		digits = append(digits, d)

		// Remove one digit from the head of the slice
		digits = digits[1:]

		// Get the product of current digits
		product := 1
		for _, d := range digits {
			product *= d
		}

		// Set new max
		if product > max {
			max = product
		}
	}

	return max
}

// LargestProductInSeries2 returns the largest product of l adjacent digits
// This method iterates over NUM as a big.Int. It is ~20x slower than the above
// method.
func LargestProductInSeries2(l int) (m int) {

	digits := make([]int, l)
	max := 0
	remainder := big.NewInt(0)
	huge := big.NewInt(0)

	// These will be used like constants
	ZERO := big.NewInt(0)
	TEN := big.NewInt(10)

	// Ingest the 1000 digit number
	huge.UnmarshalText([]byte(NUM))

	//Int.Cmp returns +1 if Int > y
	for huge.Cmp(ZERO) > 0 {

		// Pull the 1's digit off of huge using big's modulus function and put
		// it into remainder
		huge.DivMod(huge, TEN, remainder)

		// Append the digits slice with the remainder (converted to an int)
		digits = append(digits, int(remainder.Int64()))

		// Remove one digit from the head of the digits slice
		digits = digits[1:]

		// Get the product of current digits
		product := 1
		for _, d := range digits {
			product *= d
		}

		// Set new max
		if product > max {
			max = product
		}

	}

	return max
}

// Solve returns the answer for this Project Euler challenge
func Solve() int {
	return LargestProductInSeries(13)
}
