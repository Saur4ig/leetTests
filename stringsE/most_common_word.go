package stringsE

import (
	"regexp"
	"strings"
)

// 819. Most Common Word
// Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words.
// It is guaranteed there is at least one word that isn't banned, and that the answer is unique.
// Words in the list of banned words are given in lowercase, and free of punctuation.
// Words in the paragraph are not case sensitive.  The answer is in lowercase.
// link - https://leetcode.com/problems/most-common-word/
// todo: rewrite, now its to slow, with huge memory usage
func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)
	a := regexp.MustCompile(`[ !?',;.]+`)
	str := a.Split(paragraph, -1)
	var (
		res        string
		maxCounter int
	)
	values := make(map[string]int)
	ban := make(map[string]struct{})
	for _, val := range banned {
		ban[val] = struct{}{}
	}

	for _, val := range str {
		if _, ok := ban[val]; ok {
			continue
		}

		if res == "" {
			res = val
			maxCounter = 1
			values[val]++
			continue
		}

		if _, ok := values[val]; ok {
			values[val]++
			if values[val] > maxCounter {
				res = val
				maxCounter++
			}
		} else {
			values[val] = 1
		}
	}
	return res
}
