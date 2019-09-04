package arrays

// 35. Search Insert Position
// Given a sorted array and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
// You may assume no duplicates in the array.
// link - https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	first, last := 0, len(nums)-1
	for first <= last {
		mid := (first + last) / 2
		if nums[mid] == target {
			return mid
		}
		if last-first <= 2 {
			if nums[mid] < target {
				return last
			} else {
				if nums[first] < target {
					return mid
				}
				return first
			}
		}
		if nums[mid] < target {
			first = mid
		}
		if nums[mid] > target {
			last = mid
		}
	}
	return 0
}
