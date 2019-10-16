package __Two_Sum

// 167. Two Sum II - Input array is sorted
// Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.
// The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.
//
// Note:
// Your returned answers (both index1 and index2) are not zero-based.
// You may assume that each input would have exactly one solution and you may not use the same element twice.
// link - https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func twoSumSorted(numbers []int, target int) []int {
	values := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		if pos, ok := values[target-numbers[i]]; ok {
			return []int{pos + 1, i + 1}
		}
		values[numbers[i]] = i
	}
	return nil
}
