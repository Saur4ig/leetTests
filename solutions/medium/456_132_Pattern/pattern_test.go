package _56_132_Pattern

import (
	"testing"
)

func Test_find132pattern(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "Example 2",
			nums: []int{3, 1, 4, 2},
			want: true,
		},
		{
			name: "Example 3",
			nums: []int{-1, 3, 2, 0},
			want: true,
		},
		{
			name: "Example 4",
			nums: []int{3, 5, 0, 3, 4},
			want: true,
		},
		{
			name: "Example 5",
			nums: []int{1, 4, 0, -1, -2, -3, -1, -2},
			want: true,
		},
		{
			name: "Example 6",
			nums: []int{1, 0, 1, -4, -3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(tt.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
