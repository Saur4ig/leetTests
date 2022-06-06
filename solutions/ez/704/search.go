package _04

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	first, last := 0, len(nums)-1

	for first <= last {
		mid := first + (last-first)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}

	return -1
}
