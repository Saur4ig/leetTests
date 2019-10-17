package _09_Fibonacci_Number

import (
	"math"
)

// Time Complexity: O(n)
// Extra Space: O(1)
func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}
	prev1, prev2, res := 0, 1, 0
	for i := 2; i <= N; i++ {
		res = prev1 + prev2
		prev1, prev2 = prev2, res
	}
	return res
}

// Time Complexity: O(1)
// Space Complexity: O(1)
func fibFake(N int) int {
	return int(math.Round(math.Pow(math.Phi, float64(N)) / math.Sqrt(5)))
}
