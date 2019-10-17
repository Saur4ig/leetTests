package __Palindrome_Number

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
