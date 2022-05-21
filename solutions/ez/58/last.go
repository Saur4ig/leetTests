package _8

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	l := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 { // 32 = " "
			break
		}
		l++
	}
	return l
}
