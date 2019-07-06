package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "Example 1",
			input: []int{0, 1, 0},
			want:  1,
		},
		{
			name:  "Example 2",
			input: []int{0, 2, 1, 0},
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := peakIndexInMountainArray(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
