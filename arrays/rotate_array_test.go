package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  []int
	}{
		{
			name:  "Example 1",
			input: []int{1, 2, 3, 4, 5, 6, 7},
			k:     3,
			want:  []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:  "Example 2",
			input: []int{-1, -100, 3, 99},
			k:     2,
			want:  []int{3, 99, -1, -100},
		},
		{
			name:  "Example 3",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			k:     3,
			want:  []int{13, 14, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			name:  "Example 4",
			input: []int{1, 2},
			k:     3,
			want:  []int{2, 1},
		},
		{
			name:  "Example 5",
			input: []int{1, 2, 3},
			k:     4,
			want:  []int{3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.input, tt.k)
			require.Equal(t, tt.want, tt.input)
		})
	}
}
