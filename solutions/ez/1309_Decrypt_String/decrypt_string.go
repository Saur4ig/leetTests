package _309_Decrypt_String

import (
	"strconv"
	"strings"
)

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Decrypt String from Alphabet to Integer Mapping.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Decrypt String from Alphabet to Integer Mapping.
func freqAlphabets(s string) string {
	var (
		res    strings.Builder
		symbol int32
	)
	for i, char := range s {
		if string(char) != "#" {
			dig, _ := strconv.Atoi(string(char))
			if symbol == 0 {
				symbol = int32(dig)
				if len(s)-1 == i {
					res.WriteRune(symbol + 96)
				}
			} else {
				if len(s)-1 >= i+1 {
					if string(s[i+1]) == "#" {
						res.WriteRune(symbol*10 + int32(dig) + 96)
						symbol = 0
					} else {
						res.WriteRune(symbol + 96)
						symbol = int32(dig)
					}
				} else {
					res.WriteRune(symbol + 96)
					res.WriteRune(int32(dig) + 96)
				}
			}
		}
	}
	return res.String()
}
