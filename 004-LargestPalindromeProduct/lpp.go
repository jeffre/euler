package euler004

type PalindromeProduct struct {
	Palindrome int
	Products   []int
}

// Reverse reverses num efficiently by using integer division and modulus
func Reverse(num int) (reversed int) {
	for num > 0 {
		reversed *= 10
		reversed += num % 10
		num /= 10
	}
	return reversed
}

// IsPalindrome checks if num is a palindrome
func IsPalindrome(num int) bool {
	return num == Reverse(num)
}

// LargestPalindromeProduct finds the largest palindrome that is the product
// of two integers that are less than max
func LargestPalindromeProduct(max int) PalindromeProduct {
	p := PalindromeProduct{0, []int{0, 0}}

	// Iterate backwards from max for the first integer
	for i := max; i > 0; i-- {

		// Iterate backwards from i for the second integer. We
		// use i instead of max because doing so eliminates duplicative
		// work.
		for j := i; j > 0; j-- {

			ij := i * j

			// If the product of the two integers is less than the
			// current known palindrome then break out of this loop
			// because subsequent loops will only be getting
			// smaller
			if ij < p.Palindrome {
				break
			}

			// If it is a palindrome then overwrite p and then
			// break out of this loop because subsequent loops will
			// only be getting smaller
			if IsPalindrome(ij) {
				p.Palindrome = ij
				p.Products = []int{i, j}
				break
			}
		}
	}
	return p
}

func Solve() int {
	return LargestPalindromeProduct(999).Palindrome
}
