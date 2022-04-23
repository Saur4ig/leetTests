package __Longest_Substring_Without_Repeating_Characters

func lengthOfLongestSubstring(s string) int {
	max := 0
	curr := 0
	v := make(map[rune]int)

	for i, val := range s {
		if cur, ok := v[val]; ok && cur >= curr {
			curr = cur + 1
		}

		if i-curr+1 > max {
			max = i - curr + 1
		}

		v[val] = i
	}

	return max
}
