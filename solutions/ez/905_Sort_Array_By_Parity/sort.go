package _05_Sort_Array_By_Parity

func sortArrayByParity(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for _, v := range nums {
		if v%2 == 0 { // even
			res[left] = v
			left++
		} else { // odd
			res[right] = v
			right--
		}
	}
	return res
}
