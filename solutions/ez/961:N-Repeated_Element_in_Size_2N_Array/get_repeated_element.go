package _61_N_Repeated_Element_in_Size_2N_Array

func repeatedNTimes(A []int) int {
	nums := make(map[int]bool)
	for _, val := range A {
		if _, ok := nums[val]; ok {
			return val
		}
		nums[val] = true
	}
	return -1
}
