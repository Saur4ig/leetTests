package stringsE

import (
	"strings"
)

// 520. Detect Capital
// Given a word, you need to judge whether the usage of capitals in it is right or not.
//
// We define the usage of capitals in a word to be right when one of the following cases holds:
//
// All letters in this word are capitals, like "USA".
// All letters in this word are not capitals, like "leetcode".
// Only the first letter in this word is capital, like "Google".
// Otherwise, we define that this word doesn't use capitals in a right way.
// link - https://leetcode.com/problems/detect-capital/
func detectCapitalUse(word string) bool {
	if word == "" || word == " " {
		return false
	}
	if strings.ToLower(word) == word || strings.ToUpper(word) == word || strings.ToLower(word[1:]) == word[1:] {
		return true
	}
	return false
}
