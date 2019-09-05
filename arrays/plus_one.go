package arrays

// 66. Plus One
// Given a non-empty array of digits representing a non-negative integer, plus one to the integer.
// The digits are stored such that the most significant digit is at the head of the list,
// and each element in the array contain a single digit.
// You may assume the integer does not contain any leading zero, except the number 0 itself.
// link - https://leetcode.com/problems/plus-one/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		r := digits[i] + 1
		if r >= 10 {
			digits[i] = r % 10
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[i] = r
			return digits
		}
	}
	return digits
}
