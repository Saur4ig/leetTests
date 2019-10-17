package _89_Find_the_Difference

func findTheDifference(s string, t string) byte {
	first := make(map[string]int)
	for _, char := range s {
		first[string(char)]++
	}
	for _, char := range t {
		if _, ok := first[string(char)]; !ok {
			return byte(char)
		} else {
			first[string(char)]--
			if first[string(char)] < 0 {
				return byte(char)
			}
		}
	}
	return 0
}
