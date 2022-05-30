package _04

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dep := 0
	return max(root, dep)
}

func max(node *TreeNode, prevDep int) int {
	prevDep++
	r, l := prevDep, prevDep
	if node.Right != nil {
		r = max(node.Right, r)
	}
	if node.Left != nil {
		l = max(node.Left, l)
	}
	if r > l {
		return r
	}
	return l
}
