package _44_Backspace_String_Compare

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Backspace String Compare.
// Memory Usage: 2.2 MB, less than 25.00% of Go online submissions for Backspace String Compare.
func backspaceCompare(S string, T string) bool {
	return back(S) == back(T)
}

func back(v string) string {
	var res string
	for _, char := range v {
		if string(char) == "#" {
			if len(res) != 0 {
				res = res[:len(res)-1]
			}
		} else {
			res += string(char)
		}
	}
	return res
}
