package stringsE

import (
	"strings"
	"unicode"
)

// 125. Valid Palindrome
// Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
// Note: For the purpose of this problem, we define empty string as valid palindrome.
// link - https://leetcode.com/problems/valid-palindrome/
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
