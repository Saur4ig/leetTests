package _56_132_Pattern

import (
	"math"
)

// o(n^2)
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	min := nums[0]

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > min && nums[i] > nums[j] {
				return true
			}
			if nums[i] < min {
				min = nums[i]
			}
		}
	}

	return false
}

// o(n)
func find132pattern2(nums []int) bool {
	stack := make([]int, 0)
	max := math.MinInt
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < max {
			return true
		}

		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			max = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, nums[i])
	}
	return false
}
