package __string_to_int

import (
	"math"
	"strings"
)

var digits = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	if s == "" {
		return 0
	}

	var isMinus bool

	ints := make([]int, 0)

	if string(s[0]) == "-" {
		isMinus = true
		s = s[1:]
	} else if string(s[0]) == "+" {
		s = s[1:]
	}

	for _, char := range s {
		if n, ok := digits[string(char)]; ok {
			if n == 0 && len(ints) == 0 {
				continue
			}
			ints = append(ints, n)
		} else {
			break
		}
	}

	if len(ints) > 10 {
		if isMinus {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	res := 0
	counter := 1

	for i := len(ints); i > 0; i-- {
		res += ints[i-1] * counter
		counter *= 10
	}

	if res > math.MaxInt32 || res < math.MinInt32 || counter < math.MinInt32 {
		if isMinus {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	if isMinus {
		return -res
	}

	return res
}
