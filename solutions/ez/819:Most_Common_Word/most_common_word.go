package _19_Most_Common_Word

import (
	"regexp"
	"strings"
)

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
