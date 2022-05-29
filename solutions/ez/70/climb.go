package _0

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	oneBefore := 2
	twoBefore := 1
	res := 0
	for i := 2; i < n; i++ {
		res = oneBefore + twoBefore
		twoBefore = oneBefore
		oneBefore = res
	}
	return res
}
