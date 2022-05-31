package _10

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return getMaxHeight(root) != -1
}

func getMaxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := getMaxHeight(root.Left)
	r := getMaxHeight(root.Right)

	if l == -1 || r == -1 {
		return -1
	}
	if l-r > 1 || l-r < -1 {
		return -1
	}

	if l > r {
		return l + 1
	}
	return r + 1
}
