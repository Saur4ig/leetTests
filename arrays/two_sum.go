package arrays

import (
	"log"
)

// todo: do it
// 1. Two Sum
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// link - https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		f := nums[i+1:]
		log.Println("111", f)
		for j := i; j <
			len(nums[i+1:]); j++ {
			s1 := nums[i]
			s2 := nums[i+j+1]
			log.Print(s1, s2)
			if nums[i]+nums[i+j] == target {
				return []int{i, j + i}
			}
		}
	}
	return nil
}
