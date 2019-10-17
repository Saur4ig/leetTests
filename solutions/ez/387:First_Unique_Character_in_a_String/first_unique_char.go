package _87_First_Unique_Character_in_a_String

func firstUniqChar(s string) int {
	vals := make(map[string]int, len(s))
	for _, char := range s {
		vals[string(char)]++
	}

	for i, char := range s {
		if val, ok := vals[string(char)]; ok && val == 1 {
			return i
		}
	}
	return -1
}
