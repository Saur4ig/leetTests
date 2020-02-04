package _281_Subtract_sum_of_int

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Subtract the Product and Sum of Digits of an Integer.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Subtract the Product and Sum of Digits of an Integer.
func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	for n >= 1 {
		last := n % 10
		sum += last
		product *= last
		n = n / 10
	}
	return product - sum
}
