package _69_Majority_Element

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
