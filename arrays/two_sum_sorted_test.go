package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSumSorted(t *testing.T) {
	tests := []struct {
		name        string
		inputArr    []int
		inputTarget int
		want        []int
	}{
		{
			name:        "Example 1",
			inputArr:    []int{2, 7, 11, 15},
			inputTarget: 9,
			want:        []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSumSorted(tt.inputArr, tt.inputTarget)
			require.Equal(t, tt.want, got)
		})
	}
}
