package __Reverse_Integer

func strStr(haystack string, needle string) int {
	if haystack == "" || needle == "" {
		return 0
	}

	for i, val := range haystack {
		if val == rune(needle[0]) {
			skipp := false
			for j := 1; j < len(needle); j++ {
				if i+j < len(haystack) {
					if rune(needle[j]) != rune(haystack[i+j]) {
						skipp = true
						break
					}
				} else {
					skipp = true
					break
				}
			}
			if !skipp {
				return i
			}
		}
	}

	return -1
}
