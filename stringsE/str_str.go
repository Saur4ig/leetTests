package stringsE

// 28. Implement strStr()
// Implement strStr() http://www.cplusplus.com/reference/cstring/strstr/
// Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
// link - https://leetcode.com/problems/implement-strstr/
func strStr(haystack string, needle string) int {
	if (needle == "" && haystack != "") || (needle == "" && haystack == "") {
		return 0
	}
	for i, char := range haystack {
		if string(char) == string(needle[0]) {
			if len(haystack[i:]) < len(needle) {
				return -1
			}
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}
