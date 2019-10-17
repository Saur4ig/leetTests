package _44_Reverse_String

// here is the simplest solution, but not the fastest one, more better to do it like with array
func ReverseString(input string) string {
	var result string
	for _, char := range input {
		result = string(char) + result
	}
	return result
}
