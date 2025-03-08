package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))  // Output: true
	fmt.Println(isPalindrome(-121)) // Output: false
}

func isPalindrome(x int) bool {
	// Negative numbers and numbers ending in 0 (except 0 itself) are not palindromes
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversed := 0
	original := x
	for x > 0 {
		lastDigit := x % 10
		reversed = reversed*10 + lastDigit
		x /= 10
	}
	return original == reversed
}
