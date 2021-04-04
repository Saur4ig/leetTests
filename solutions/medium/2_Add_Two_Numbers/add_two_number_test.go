package __Add_Two_Numbers

import (
	"reflect"
	"testing"

	"github.com/Saur4ig/leetTests/list"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *list.ListNode
		l2   *list.ListNode
		want *list.ListNode
	}{
		{
			name: "Example 1",
			l1: &list.ListNode{
				Val: 2,
				Next: &list.ListNode{
					Val: 4,
					Next: &list.ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			l2: &list.ListNode{
				Val: 5,
				Next: &list.ListNode{
					Val: 6,
					Next: &list.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			want: &list.ListNode{
				Val: 7,
				Next: &list.ListNode{
					Val: 0,
					Next: &list.ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "Example 2",
			l1: &list.ListNode{
				Val: 9,
				Next: &list.ListNode{
					Val: 9,
					Next: &list.ListNode{
						Val: 9,
						Next: &list.ListNode{
							Val: 9,
							Next: &list.ListNode{
								Val: 9,
								Next: &list.ListNode{
									Val: 9,
									Next: &list.ListNode{
										Val:  9,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
			l2: &list.ListNode{
				Val: 9,
				Next: &list.ListNode{
					Val: 9,
					Next: &list.ListNode{
						Val: 9,
						Next: &list.ListNode{
							Val:  9,
							Next: nil,
						},
					},
				},
			},
			want: &list.ListNode{
				Val: 8,
				Next: &list.ListNode{
					Val: 9,
					Next: &list.ListNode{
						Val: 9,
						Next: &list.ListNode{
							Val: 9,
							Next: &list.ListNode{
								Val: 0,
								Next: &list.ListNode{
									Val: 0,
									Next: &list.ListNode{
										Val: 0,
										Next: &list.ListNode{
											Val:  1,
											Next: nil,
										},
									},
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
			if got := addTwoNumbers(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
