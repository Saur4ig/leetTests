package _01

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return rlSymmetric(root.Left, root.Right)
}

func rlSymmetric(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return rlSymmetric(left.Left, right.Right) && rlSymmetric(left.Right, right.Left)
}
