package _57_Reverse_Words_in_a_String_III

import (
	"strings"
	"unicode"
)

// Runtime: 36 ms, faster than 19.23% of Go online submissions for Reverse Words in a String III.
// Memory Usage: 6.6 MB, less than 100.00% of Go online submissions for Reverse Words in a String III.
func reverseWords(s string) string {
	var (
		res  strings.Builder
		word string
	)
	for _, val := range s {
		if unicode.IsSpace(val) {
			res.WriteString(word)
			res.WriteString(" ")
			word = ""
			continue
		}
		word = string(val) + word
	}
	res.WriteString(word)
	return res.String()
}
