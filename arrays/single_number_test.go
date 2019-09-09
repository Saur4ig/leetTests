package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Example 1",
			input: []int{2, 2, 1},
			want:  1,
		},
		{
			name:  "Example 2",
			input: []int{4, 1, 2, 1, 2},
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
