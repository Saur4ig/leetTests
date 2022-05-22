package _9

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	h := x
	for h*h > x {
		h = (h + x/h) / 2
	}
	return h
}
