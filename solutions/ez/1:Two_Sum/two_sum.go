package __Two_Sum

// very slow
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i+1:]); j++ {
			if nums[i]+nums[i+j+1] == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return nil
}

// it uses more memory, but more faster
func twoSumFaster(nums []int, target int) []int {
	values := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if pos, ok := values[target-nums[i]]; ok {
			return []int{pos, i}
		}
		values[nums[i]] = i
	}
	return nil
}
