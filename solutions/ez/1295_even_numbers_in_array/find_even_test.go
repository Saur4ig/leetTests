package _295_even_numbers_in_array

import (
	"testing"
)

func Test_findNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{12, 345, 2, 6, 7896},
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{555, 901, 482, 1771},
			want: 1,
		},
		{
			name: "Example 3",
			nums: []int{},
			want: 0,
		},
		{
			name: "Example 4",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 0,
		},
		{
			name: "Example 5",
			nums: []int{11, 12, 13, 14, 15, 16, 9999},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.nums); got != tt.want {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
