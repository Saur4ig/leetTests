package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Example 1",
			input: []int{1, 1, 0, 1, 1, 1},
			want:  3,
		},
		{
			name:  "Example 2",
			input: []int{1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1},
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxConsecutiveOnes(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
