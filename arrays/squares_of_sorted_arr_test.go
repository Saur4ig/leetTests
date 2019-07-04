package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Example 1",
			input: []int{-4, -1, 0, 3, 10},
			want:  []int{0, 1, 9, 16, 100},
		},
		{
			name:  "Example 2",
			input: []int{-7, -3, 2, 3, 11},
			want:  []int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedSquares(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
