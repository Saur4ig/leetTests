package _25_Valid_Palindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	if s == "" || s == " " {
		return true
	}
	if len(s) == 1 {
		return true
	}
	r := []rune(s)
	var left, right = 0, len(r) - 1
	for i := 0; i < len(r)-1; i++ {
		if !unicode.IsLetter(r[left]) && !unicode.IsNumber(r[left]) {
			left++
			continue
		}
		if !unicode.IsLetter(r[right]) && !unicode.IsNumber(r[right]) {
			right--
			continue
		}
		if strings.ToLower(string(r[left])) != strings.ToLower(string(r[right])) {
			return false
		}
		left++
		right--
		if left == right {
			return true
		}
	}
	return true
}
