package arrays

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
			inputArr:    []int{1, 2, 3, 3},
			inputTarget: 9,
			want:        []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.inputArr, tt.inputTarget)
			require.Equal(t, tt.want, got)
		})
	}
}
