package __Reverse_Integer

import (
	"math"
)

func reverse(x int) int {
	var (
		res int
		neg bool
	)
	if x < 0 {
		x = int(math.Abs(float64(x)))
		neg = true
	}
	for x > 0 {
		last := x % 10
		res = (res * 10) + last
		x /= 10
	}
	if math.MinInt32 > res || math.MaxInt32 < res {
		return 0
	}
	if neg {
		return 0 - res
	}
	return res
}
