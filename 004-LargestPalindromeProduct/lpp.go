package euler004

type PalindromeProduct struct {
	Palindrome int
	Products   []int
}

// Reverse reverses num effeciently by using integer division and modulus
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

	for i := max; i > 0; i-- {
		for j := max; j > 0; j-- {
			ij := i * j
			if p.Palindrome > ij {
				break
			}
			if IsPalindrome(ij) {
				p.Palindrome = ij
				p.Products = []int{i, j}
				break
			}
		}
	}
	return p
}
