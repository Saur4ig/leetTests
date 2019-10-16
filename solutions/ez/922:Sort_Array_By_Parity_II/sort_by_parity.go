package _22_Sort_Array_By_Parity_II

func sortArrayByParityII(A []int) []int {
	nextE, nextO := 0, 1
	arr := make([]int, len(A))
	for _, val := range A {
		if val%2 == 0 {
			arr[nextE] = val
			nextE += 2
			continue
		}
		arr[nextO] = val
		nextO += 2
	}
	return arr
}
