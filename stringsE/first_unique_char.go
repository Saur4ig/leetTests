package stringsE

// 387. First Unique Character in a String
// Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
// link - https://leetcode.com/problems/first-unique-character-in-a-string/
func firstUniqChar(s string) int {
	vals := make(map[string]int, len(s))
	for _, char := range s {
		vals[string(char)]++
	}

	for i, char := range s {
		if val, ok := vals[string(char)]; ok && val == 1 {
			return i
		}
	}
	return -1
}
