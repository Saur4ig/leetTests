package _36_Single_Number

func singleNumber(nums []int) int {
	counter := make(map[int]int)
	for _, val := range nums {
		counter[val]++
		if counter[val]%2 == 0 {
			delete(counter, val)
		}
	}
	for key := range counter {
		return key
	}
	return 0
}
