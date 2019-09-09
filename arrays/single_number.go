package arrays

// 136. Single Number
// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
//
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
// link - https://leetcode.com/problems/single-number/
func singleNumber(nums []int) int {
	counter := make(map[int]int)
	for _, val := range nums {
		counter[val]++
		if counter[val]%2 == 0 {
			delete(counter, val)
		}
	}
	for key := range counter {
		return key
	}
	return 0
}
