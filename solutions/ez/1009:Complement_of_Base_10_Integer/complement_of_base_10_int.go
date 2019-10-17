package _009_Complement_of_Base_10_Integer

import (
	"strconv"
	"strings"
)

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
