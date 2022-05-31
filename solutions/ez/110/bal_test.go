package _10

import (
	"testing"
)

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 4,
							},
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			want: false,
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
