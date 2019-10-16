package _6_Plus_One

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		r := digits[i] + 1
		if r >= 10 {
			digits[i] = r % 10
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[i] = r
			return digits
		}
	}
	return digits
}
