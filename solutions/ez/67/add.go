package _7

func addBinary(a string, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	add := false
	if len(a) < len(b) {
		a, b = b, a
	}

	max := len(a) - 1
	min := len(b) - 1

	res := ""
	for i := 0; i < len(a); i++ {
		if a[max-i] == 49 && b[min-i] == 49 {
			if add {
				res = "1" + res
			} else {
				res = "0" + res
			}
			add = true
		} else if (a[max-i] == 48 && b[min-i] == 49) || (a[max-i] == 49 && b[min-i] == 48) {
			if add {
				res = "0" + res
				add = true
			} else {
				res = "1" + res
			}
		} else {
			if add {
				res = "1" + res
				add = false
			} else {
				res = "0" + res
			}
		}
		if i >= min {
			if add {
				return addBinary(a[:max-i], "1") + res
			}
			return a[:max-i] + res
		}
	}

	return res
}
