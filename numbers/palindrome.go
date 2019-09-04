package numbers

// 9. Palindrome Number
// Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
// link - https://leetcode.com/problems/palindrome-number
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	a, res := x, 0
	for x > 0 {
		last := x % 10
		res = (res * 10) + last
		x /= 10
	}
	return a == res
}
