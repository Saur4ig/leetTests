package _7_Letter_Combinations

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	buttons := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	if len(digits) == 1 {
		return buttons[string(digits[0])]
	}

	res := buttons[string(digits[0])]

	for i := 1; i < len(digits); i++ {
		res = addOne(res, buttons[string(digits[i])])
	}

	return res
}

func addOne(prev, new []string) []string {
	n := make([]string, 0)
	for _, val := range prev {
		for _, s := range new {
			n = append(n, val+s)
		}
	}
	return n
}
