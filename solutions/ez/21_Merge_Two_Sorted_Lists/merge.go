package _1_Merge_Two_Sorted_Lists

import (
	"github.com/Saur4ig/leetTests/list"
)

func mergeTwoLists(l1 *list.ListNode, l2 *list.ListNode) *list.ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	res := &list.ListNode{}
	current := res
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}

	if l2 != nil {
		current.Next = l2
	}
	return res.Next
}
