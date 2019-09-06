package arrays

// 169. Majority Element
// Given an array of size n, find the majority element.
// The majority element is the element that appears more than ⌊ n/2 ⌋ times.
// You may assume that the array is non-empty and the majority element always exist in the array.
//
// link - https://leetcode.com/problems/majority-element/
// todo: create solution without map
func majorityElement(nums []int) int {
	pretender := nums[0]
	values := make(map[int]int)
	for _, val := range nums {
		values[val]++
		if val == pretender {
			if values[val] > len(nums)/2 {
				return val
			}
		} else {
			if values[val] > values[pretender] {
				pretender = val
			}
		}
	}
	return pretender
}
