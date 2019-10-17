package _20_Detect_Capital

import (
	"strings"
)

func detectCapitalUse(word string) bool {
	if word == "" || word == " " {
		return false
	}
	if strings.ToLower(word) == word || strings.ToUpper(word) == word || strings.ToLower(word[1:]) == word[1:] {
		return true
	}
	return false
}
