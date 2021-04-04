package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsEqual(first, second *ListNode) bool {
	for first != nil {
		if !isNodeEq(first, second) {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
}

func isNodeEq(first, second *ListNode) bool {
	if (first.Val != second.Val) || (first.Next != nil && second.Next == nil) || (first.Next == nil && second.Next != nil) {
		return false
	}
	return true
}
