package stringsE

// 942. DI String Match
// Given a string S that only contains "I" (increase) or "D" (decrease), let N = S.length.
// Return any permutation A of [0, 1, ..., N] such that for all i = 0, ..., N-1:

// Runtime: 160 ms, faster than 33.54% of Go online submissions for DI String Match.
// Memory Usage: 7.9 MB, less than 100.00% of Go online submissions for DI String Match.
// link - https://leetcode.com/problems/di-string-match/
func diStringMatch(S string) []int {
	if len(S) == 0 {
		return nil
	}
	res := make([]int, len(S)+1)
	min, max := 0, len(S)
	for i := 0; i < len(S); i++ {
		if string(S[i]) == "I" {
			res[i] = min
			min++
		} else {
			res[i] = max
			max--
		}
	}
	res[len(S)] = min
	return res
}
