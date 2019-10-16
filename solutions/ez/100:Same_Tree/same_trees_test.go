package _00_Same_Tree

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Saur4ig/leetTests/trees"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name   string
		first  *trees.TreeNode
		second *trees.TreeNode
		want   bool
	}{
		{
			name: "Example 1",
			first: &trees.TreeNode{
				Val: 1,
				Left: &trees.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &trees.TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			second: &trees.TreeNode{
				Val: 1,
				Left: &trees.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &trees.TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			want: true,
		},
		{
			name: "Example 2",
			first: &trees.TreeNode{
				Val: 1,
				Left: &trees.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			second: &trees.TreeNode{
				Val:  1,
				Left: nil,
				Right: &trees.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			want: false,
		},
		{
			name: "Example 3",
			first: &trees.TreeNode{
				Val: 1,
				Left: &trees.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &trees.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			second: &trees.TreeNode{
				Val: 1,
				Left: &trees.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &trees.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSameTree(tt.first, tt.second)
			require.Equal(t, tt.want, got)
		})
	}
}
