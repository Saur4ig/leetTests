package _5_Search_Insert_Position

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
