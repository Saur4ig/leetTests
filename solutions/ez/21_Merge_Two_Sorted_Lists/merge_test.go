package _1_Merge_Two_Sorted_Lists

import (
	"reflect"
	"testing"

	"github.com/Saur4ig/leetTests/list"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name string
		l1   *list.ListNode
		l2   *list.ListNode
		want *list.ListNode
	}{
		{
			name: "Example 1",
			l1: &list.ListNode{
				Val: 1,
				Next: &list.ListNode{
					Val: 2,
					Next: &list.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			l2: &list.ListNode{
				Val: 1,
				Next: &list.ListNode{
					Val: 3,
					Next: &list.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			want: &list.ListNode{
				Val: 1,
				Next: &list.ListNode{
					Val: 1,
					Next: &list.ListNode{
						Val: 2,
						Next: &list.ListNode{
							Val: 3,
							Next: &list.ListNode{
								Val: 4,
								Next: &list.ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
