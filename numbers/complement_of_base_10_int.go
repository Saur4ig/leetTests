package numbers

import (
	"strconv"
	"strings"
)

// 1009. Complement of Base 10 Integer
// Every non-negative integer N has a binary representation.
// For example, 5 can be represented as "101" in binary, 11 as "1011" in binary, and so on.
// Note that except for N = 0, there are no leading zeroes in any binary representation.
// The complement of a binary representation is the number in binary you get when changing every 1 to a 0 and 0 to a 1.
// For example, the complement of "101" in binary is "010" in binary.
// For a given number N in base-10, return the complement of it's binary representation as a base-10 integer.
// link - https://leetcode.com/problems/complement-of-base-10-integer/
func bitwiseComplement(N int) int {
	dem2 := strconv.FormatInt(int64(N), 2)
	var sb strings.Builder
	for _, val := range dem2 {
		if string(val) == "1" {
			sb.WriteString("0")
		} else {
			sb.WriteString("1")
		}
	}
	i, err := strconv.ParseInt(sb.String(), 2, 64)
	if err != nil {
		return 0
	}
	return int(i)
}
