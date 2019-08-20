package stringsE

// todo: refactor this, coz pretty hard to understand
// 14. Longest Common Prefix
// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
// All given inputs are in lowercase letters a-z.
// link - https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs[0] == "" {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]

OUTER:
	for i := 1; i < len(strs); i++ {
		if strs[i] == "" {
			return ""
		}
		m := len(strs[i])
		if len(prefix) <= m {
			m = len(prefix)
		} else {
			prefix = prefix[:m]
		}
		for j := 0; j < m; j++ {
			if strs[i][j] != prefix[j] {
				if j == 0 {
					return ""
				} else {
					prefix = prefix[:j]
				}
				continue OUTER
			}
		}
	}

	return prefix
}
