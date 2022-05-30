package _9

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}

	sign := 1
	if !(divisor > 0 && dividend > 0) && !(divisor < 0 && dividend < 0) {
		sign = -1
	}
	if divisor < 0 {
		divisor *= -1
	}
	if dividend < 0 {
		dividend *= -1
	}

	if divisor == 1 {
		return res(dividend, sign)
	}
	if divisor == -1 {
		return -res(dividend, sign)
	}

	n := 0
	for dividend >= divisor {
		dividend -= divisor
		n++
	}

	return res(n, sign)
}

func res(a, sign int) int {
	if a >= math.MaxInt32 {
		if sign == -1 {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	return a * sign
}
