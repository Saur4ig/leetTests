package arrays

// 189. Rotate Array
// Given an array, rotate the array to the right by k steps, where k is non-negative.
// link - https://leetcode.com/problems/rotate-array/
func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	k = len(nums) - k%len(nums)
	copy(nums, append(nums[k:], nums[:k]...))
}
