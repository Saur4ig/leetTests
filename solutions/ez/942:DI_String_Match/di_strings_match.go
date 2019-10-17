package _42_DI_String_Match

// Runtime: 160 ms, faster than 33.54% of Go online submissions for DI String Match.
// Memory Usage: 7.9 MB, less than 100.00% of Go online submissions for DI String Match.
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
