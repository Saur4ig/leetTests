package _58_Add_Digits

func AddDigits(num int) int {
	if num < 10 {
		return num
	}
	for num >= 10 {
		num = num%10 + (num / 10)
	}
	return num
}
