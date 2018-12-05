package stringsE

func ReverseString(input string) string {
	var result string
	for _, char := range input {
		result = string(char) + result
	}
	return result
}
