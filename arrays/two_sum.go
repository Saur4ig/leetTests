package arrays

// todo: do it
// 1. Two Sum
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// link - https://leetcode.com/problems/two-sum/
// very slow
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i+1:]); j++ {
			if nums[i]+nums[i+j+1] == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return nil
}

// it uses more memory, but more faster
func twoSumFaster(nums []int, target int) []int {
	values := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if pos, ok := values[target-nums[i]]; ok {
			return []int{pos, i}
		}
		values[nums[i]] = i
	}
	return nil
}
