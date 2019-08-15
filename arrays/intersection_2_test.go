package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name   string
		input1 []int
		input2 []int
		want   []int
	}{
		{
			name:   "Example 1",
			input1: []int{1, 2, 2, 1},
			input2: []int{2, 2},
			want:   []int{2, 2},
		},
		{
			name:   "Example 2",
			input1: []int{4, 9, 5},
			input2: []int{9, 4, 9, 8, 4},
			want:   []int{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect(tt.input1, tt.input2)
			require.Equal(t, tt.want, got)
		})
	}
}
