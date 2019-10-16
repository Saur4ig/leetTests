package _5_Search_Insert_Position

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "Example 2",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			name:   "Example 3",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			name:   "Example 4",
			nums:   []int{1, 3, 5, 6},
			target: 0,
			want:   0,
		},
		{
			name:   "Example 5",
			nums:   []int{1},
			target: 2,
			want:   1,
		},
		{
			name:   "Example 6",
			nums:   []int{1, 2},
			target: 0,
			want:   0,
		},
		{
			name:   "Example 7",
			nums:   []int{1, 2},
			target: 3,
			want:   2,
		},
		{
			name:   "Example 8",
			nums:   []int{},
			target: 5,
			want:   0,
		},
		{
			name:   "Example 9",
			nums:   []int{1, 2},
			target: 10,
			want:   2,
		},
		{
			name:   "Example 10",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15, 16, 17},
			target: 13,
			want:   12,
		},
		{
			name:   "Example 11",
			nums:   []int{0, 2, 5, 6, 10, 11, 25},
			target: 15,
			want:   6,
		},
		{
			name:   "Example 12",
			nums:   []int{123, 127, 150, 190, 250},
			target: 177,
			want:   3,
		},
		{
			name:   "Example 13",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 7,
			want:   6,
		},
		{
			name:   "Example 14",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 1,
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			require.Equal(t, tt.want, got)
		})
	}
}
