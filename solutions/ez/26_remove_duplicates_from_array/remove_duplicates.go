package __Reverse_Integer

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	prevValue := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == prevValue {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			prevValue = nums[i]
		}
	}

	return len(nums)
}
