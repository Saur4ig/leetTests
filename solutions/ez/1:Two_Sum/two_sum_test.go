package __Two_Sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
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
			want:        []int{0, 1},
		},
		{
			name:        "Example 2",
			inputArr:    []int{3, 2, 4},
			inputTarget: 6,
			want:        []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.inputArr, tt.inputTarget)
			require.Equal(t, tt.want, got)
		})
	}
}
