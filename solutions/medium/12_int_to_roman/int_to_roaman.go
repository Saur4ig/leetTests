package _2_int_to_roman

import (
	"fmt"
)

//Runtime: 24 ms, faster than 10.06% of Go online submissions for Integer to Roman.
//Memory Usage: 4.3 MB, less than 15.08% of Go online submissions for Integer to Roman.
func intToRoman(num int) string {
	dictionary := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	res := ""
	counter := 1

	for num > 0 {
		last := num % 10
		if last != 0 {

			n := last * counter
			if s, ok := dictionary[n]; ok {
				res = fmt.Sprintf("%s%s", s, res)
			} else {
				closest, multiplier, num := getClosestNMult(last)
				if multiplier {
					for i := 0; i < num; i++ {
						res = fmt.Sprintf("%s%s", dictionary[1*counter], res)
					}
					if closest != 1 {
						res = fmt.Sprintf("%s%s", dictionary[closest*counter], res)
					}
				} else {
					res = fmt.Sprintf("%s%s%s", dictionary[num*counter], dictionary[closest*counter], res)
				}
			}
		}

		counter *= 10
		num /= 10
	}

	return res
}

func getClosestNMult(n int) (int, bool, int) {
	if n > 1 && n < 5 {
		if n == 4 {
			return 5, false, 1
		} else {
			return 1, true, n
		}
	}
	if n == 9 {
		return 10, false, 1
	}

	return 5, true, n - 5
}
