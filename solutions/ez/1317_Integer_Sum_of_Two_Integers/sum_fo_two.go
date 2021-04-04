package _317_Integer_Sum_of_Two_Integers

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Convert Integer to the Sum of Two No-Zero Integers.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Convert Integer to the Sum of Two No-Zero Integers.
// todo: pretty messy, refactoring needed
func getNoZeroIntegers(n int) []int {
	if n <= 10 {
		return []int{1, n - 1}
	}
	period, left := 1, 0
	full := n
	for n >= 10 {
		last := n % 10
		switch last {
		case 1:
			left += 2 * period
		default:
			left += 1 * period
		}
		rr := full - left
		if !isZeroIn(rr) {
			return []int{left, rr}
		} else {
			n = rr / period
		}
		period *= 10
	}

	return []int{-1, -1}
}

func isZeroIn(n int) bool {
	for n >= 10 {
		if n%10 == 0 {
			return true
		}
		n = n / 10
	}
	return false
}
