package _08

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}

	return add(nums, 0, len(nums)-1)
}

func add(n []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	return &TreeNode{
		Val:   n[(left+right)/2],
		Left:  add(n, left, (left+right)/2-1),
		Right: add(n, (left+right)/2+1, right),
	}
}
