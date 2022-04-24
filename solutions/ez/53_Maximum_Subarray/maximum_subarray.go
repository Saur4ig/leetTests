package __Reverse_Integer

import (
	"math"
)

// O(N^2)
func maxSubArray(nums []int) int {
	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		currentMax := 0
		for j := i; j < len(nums); j++ {
			currentMax += nums[j]
			if currentMax > max {
				max = currentMax
			}
		}
	}
	return max
}

// O(N)
func maxSubArray2(nums []int) int {
	max := math.MinInt
	curr := 0

	for _, v := range nums {
		curr += v
		if curr > max {
			max = curr
		}
		if curr < 0 {
			curr = 0
		}
	}

	return max
}
