package arrays

// 485. Max Consecutive Ones
// Given a binary array, find the maximum number of consecutive 1s in this array.
// link - https://leetcode.com/problems/max-consecutive-ones/
func findMaxConsecutiveOnes(nums []int) int {
	var pretender, max int
	for _, val := range nums {
		if val == 0 {
			pretender = 0
		} else {
			pretender++
			if pretender > max {
				max = pretender
			}
		}
	}
	return max
}
