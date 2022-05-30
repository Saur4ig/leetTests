package _18

func maxProduct(words []string) int {
	if len(words) <= 1 {
		return 0
	}

	max := 0
	for i := 0; i < len(words); i++ {
		for j := i; j < len(words); j++ {
			if isUnique(words[i], words[j]) {
				n := len(words[i]) * len(words[j])
				if n > max {
					max = n
				}
			}
		}
	}
	return max
}

func isUnique(a, b string) bool {
	for _, aa := range a {
		for _, bb := range b {
			if bb == aa {
				return false
			}
		}
	}
	return true
}
