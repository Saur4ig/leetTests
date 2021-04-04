package _3_remove_duplicates_from_sorted_list

import (
	"testing"

	"github.com/Saur4ig/leetTests/list"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		head *list.ListNode
		want *list.ListNode
	}{
		{
			name: "Example 1",
			head: &list.ListNode{
				Val: 1,
				Next: &list.ListNode{
					Val: 1,
					Next: &list.ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: &list.ListNode{
				Val: 1,
				Next: &list.ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.head); !list.IsEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
