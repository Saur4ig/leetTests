package numbers

// Add Digits 258
// Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.
func AddDigits(num int) int {
	if num < 10 {
		return num
	}
	for num >= 10 {
		num = num%10 + (num / 10)
	}
	return num
}
