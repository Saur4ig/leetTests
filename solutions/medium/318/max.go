package _18

func maxProduct(words []string) int {
	if len(words) <= 1 {
		return 0
	}

	max := 0
	for i := 0; i < len(words); i++ {
		for j := i; j < len(words); j++ {
			n := getN(words[i], words[j])
			if n > max {
				max = n
			}

		}
	}
	return max
}

func getN(a, b string) int {
	for _, aa := range a {
		for _, bb := range b {
			if bb == aa {
				return 0
			}
		}
	}
	return len(a) * len(b)
}
