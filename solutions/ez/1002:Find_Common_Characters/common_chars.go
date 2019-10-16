package _002_Find_Common_Characters

import (
	"strings"
)

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
