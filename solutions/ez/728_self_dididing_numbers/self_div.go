package _28_self_dididing_numbers

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Self Dividing Numbers.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Self Dividing Numbers.
func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)
	for i := left; i <= right; i++ {
		if isSelfDivNum(i) {
			res = append(res, i)
		}
	}
	return res
}

func isSelfDivNum(n int) bool {
	if n < 10 {
		return true
	}
	s := n
	for s > 0 {
		last := s % 10
		if last == 0 || n%last != 0 {
			return false
		}
		s = s / 10
	}
	return true
}
