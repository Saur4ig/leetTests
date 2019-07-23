package arrays

// 1002. Find Common Characters
// Given an array A of strings made only from lowercase letters, return a list of all characters that show up in all
// strings within the list (including duplicates).
// For example, if a character occurs 3 times in all strings but not 4 times, you need to include
// that character three times in the final answer.
// link - https://leetcode.com/problems/find-common-characters/
func commonChars(A []string) []string {
	baseMult := 10
	letters := make(map[string]int)
	res := make([]string, 0)
	for _, word := range A {
		for _, char := range word {

		}
	}
	return res
}
