package arrays

import (
	"strings"
)

// 1002. Find Common Characters
// Given an array A of strings made only from lowercase letters, return a list of all characters that show up in all
// strings within the list (including duplicates).
// For example, if a character occurs 3 times in all strings but not 4 times, you need to include
// that character three times in the final answer.
// link - https://leetcode.com/problems/find-common-characters/
// pretty average solution
func commonChars(A []string) []string {
	if len(A) == 0 {
		return []string{}
	}
	length := len(A)
	first := strings.Split(A[0], "")
	letters := make(map[string][]int)
	for _, s := range first {
		if _, ok := letters[s]; ok {
			letters[s][0]++
		} else {
			arr := make([]int, length)
			arr[0] = 1
			letters[s] = arr
		}
	}

	for i := 1; i < len(A); i++ {
		for _, char := range A[i] {
			if _, ok := letters[string(char)]; ok {
				letters[string(char)][i]++
			}
		}
	}

	res := make([]string, 0)
OUTER:
	for let, counter := range letters {
		min := counter[0]
		for _, num := range counter {
			if num == 0 {
				continue OUTER
			}
			if num < min {
				min = num
			}
		}
		for min > 0 {
			res = append(res, let)
			min--
		}
	}
	return res
}
