package __Add_Two_Numbers

import (
	"github.com/Saur4ig/leetTests/list"
)

func addTwoNumbers(l1 *list.ListNode, l2 *list.ListNode) *list.ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	res := &list.ListNode{}
	current := res
	var add int

	for l1 != nil && l2 != nil {
		current.Next = &list.ListNode{}
		current = current.Next
		temp := l1.Val + l2.Val + add
		add = 0
		if temp >= 10 {
			add = 1
			temp = temp % 10
		}
		current.Val = temp
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		if add == 0 {
			current.Next = l1
		} else {
			for l1 != nil {
				current.Next = &list.ListNode{}
				current = current.Next
				temp := l1.Val + add
				add = 0
				if temp >= 10 {
					add = 1
					temp = temp % 10
				}
				current.Val = temp
				l1 = l1.Next
			}
		}
	}

	if l2 != nil {
		if add == 0 {
			current.Next = l2
		} else {
			for l2 != nil {
				current.Next = &list.ListNode{}
				current = current.Next
				temp := l2.Val + add
				add = 0
				if temp >= 10 {
					add = 1
					temp = temp % 10
				}
				current.Val = temp
				l2 = l2.Next
			}
		}
	}

	if add != 0 {
		current.Next = &list.ListNode{Val: add}
	}

	return res.Next
}
