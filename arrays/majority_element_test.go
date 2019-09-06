package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		/*{
			name:  "Example 1",
			input: []int{3, 2, 3},
			want:  3,
		},
		{
			name:  "Example 2",
			input: []int{2, 2, 1, 1, 1, 2, 2},
			want:  2,
		},
		{
			name:  "Example 3",
			input: []int{1},
			want:  1,
		},
		{
			name:  "Example 4",
			input: []int{0, 0},
			want:  0,
		},*/
		{
			name:  "Example 5",
			input: []int{8, 8, 7, 7, 7},
			want:  7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
