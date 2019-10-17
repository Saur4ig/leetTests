package _8_strStr

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
