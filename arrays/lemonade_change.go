package arrays

// 860. Lemonade Change
// At a lemonade stand, each lemonade costs $5.
// Customers are standing in a queue to buy from you, and order one at a time (in the order specified by bills).
// Each customer will only buy one lemonade and pay with either a $5, $10, or $20 bill.  You must provide the correct change to each customer, so that the net transaction is that the customer pays $5.
// Note that you don't have any change in hand at first.
// Return true if and only if you can provide every customer with correct change.
//
// link - https://leetcode.com/problems/lemonade-change/
//
// Runtime: 28 ms, faster than 5.06% of Go online submissions for Lemonade Change.
// Memory Usage: 5.8 MB, less than 100.00% of Go online submissions for Lemonade Change.
// todo: refactoring needed.. coz too slow
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
