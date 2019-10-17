package _3_Roman_to_Integer

// pretty slow solution
func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var (
		sum      int
		skipNext bool
	)
	for i := range s {
		if skipNext {
			skipNext = false
			continue
		}
		if i == len(s)-1 || roman[string(s[i])] >= roman[string(s[i+1])] {
			sum += roman[string(s[i])]
		} else {
			sum += roman[string(s[i+1])] - roman[string(s[i])]
			skipNext = true
		}
	}
	return sum
}
