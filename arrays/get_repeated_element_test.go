package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepeatedNTimes(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Example 1",
			input: []int{1, 2, 3, 3},
			want:  3,
		},
		{
			name:  "Example 2",
			input: []int{2, 1, 2, 5, 3, 2},
			want:  2,
		},
		{
			name:  "Example 3",
			input: []int{5, 1, 5, 2, 5, 3, 5, 4},
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := repeatedNTimes(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
