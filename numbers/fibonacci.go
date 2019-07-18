package numbers

// 509. Fibonacci Number
// The Fibonacci numbers, commonly denoted F(n) form a sequence,
// called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1.
// link - https://leetcode.com/problems/fibonacci-number/
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
