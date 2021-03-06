package _60_Lemonade_Change

// Runtime: 28 ms, faster than 5.06% of Go online submissions for Lemonade Change.
// Memory Usage: 5.8 MB, less than 100.00% of Go online submissions for Lemonade Change.
func lemonadeChange(bills []int) bool {
	if len(bills) == 0 || bills[0] > 5 {
		return false
	}
	wallet := make(map[int]int)
	for _, bill := range bills {
		if bill%5 != 0 {
			return false
		}
		switch bill {
		case 10:
			if wallet[5] < 1 {
				return false
			}
			wallet[5]--
		case 20:
			if wallet[10] == 0 {
				if wallet[5] < 3 {
					return false
				} else {
					wallet[5] -= 3
				}
			} else {
				if wallet[5] == 0 {
					return false
				} else {
					wallet[10]--
					wallet[5]--
				}
			}
		case 5:
			wallet[bill]++
			continue
		default:
			return false
		}
		wallet[bill]++
	}
	return true
}

// a little bit faster
// Runtime: 16 ms, faster than 75.95% of Go online submissions for Lemonade Change.
// Memory Usage: 5.8 MB, less than 100.00% of Go online submissions for Lemonade Change.
func lemonadeChangeFaster(bills []int) bool {
	if len(bills) == 0 || bills[0] != 5 {
		return false
	}
	var bill5, bill10 int
	for _, bill := range bills {
		switch bill {
		case 5:
			bill5++
		case 10:
			if bill5 < 1 {
				return false
			}
			bill5--
			bill10++
		case 20:
			if bill10 < 1 {
				if bill5 < 3 {
					return false
				}
				bill5 -= 3
			} else {
				if bill5 < 1 {
					return false
				}
				bill5--
				bill10--
			}
		}
	}
	return true
}
