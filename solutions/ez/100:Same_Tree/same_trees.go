package _00_Same_Tree

import (
	"github.com/Saur4ig/leetTests/trees"
)

func isSameTree(p *trees.TreeNode, q *trees.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) || (p.Val != q.Val) {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
