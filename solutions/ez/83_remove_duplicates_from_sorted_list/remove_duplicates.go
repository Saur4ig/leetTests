package _3_remove_duplicates_from_sorted_list

import (
	"github.com/Saur4ig/leetTests/list"
)

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
// Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
func deleteDuplicates(head *list.ListNode) *list.ListNode {
	nl := head
	for nl != nil {
		if nl.Next != nil {
			if nl.Val == nl.Next.Val {
				nl.Next = nl.Next.Next
				continue
			}
		}
		nl = nl.Next
	}
	return head
}
