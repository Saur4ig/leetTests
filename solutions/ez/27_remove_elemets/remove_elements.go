package __Reverse_Integer

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}
