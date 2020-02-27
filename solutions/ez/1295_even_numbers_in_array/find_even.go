package _295_even_numbers_in_array

// Runtime: 4 ms, faster than 91.16% of Go online submissions for Find Numbers with Even Number of Digits.
// Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Find Numbers with Even Number of Digits.
func findNumbers(nums []int) int {
	var res int
	for _, num := range nums {
		if count(num)%2 == 0 {
			res++
		}
	}
	return res
}

func count(num int) int {
	var c int
	for num > 0 {
		c++
		num = num / 10
	}
	return c
}
