package _85_Max_Consecutive_Ones

func findMaxConsecutiveOnes(nums []int) int {
	var pretender, max int
	for _, val := range nums {
		if val == 0 {
			pretender = 0
		} else {
			pretender++
			if pretender > max {
				max = pretender
			}
		}
	}
	return max
}
