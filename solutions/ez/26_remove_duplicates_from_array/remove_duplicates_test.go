package __Reverse_Integer

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4, 5, 6, 6},
			want: 6,
		},
		{
			name: "Example 2",
			nums: []int{1, 1, 1, 1, 1, 1, 1, 1},
			want: 1,
		},
		{
			name: "Example 3",
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			name: "Example 4",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
