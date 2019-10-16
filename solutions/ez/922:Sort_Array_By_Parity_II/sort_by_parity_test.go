package _22_Sort_Array_By_Parity_II

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortArrayByParityII(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Example 1",
			input: []int{4, 2, 5, 7},
			want:  []int{4, 5, 2, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArrayByParityII(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
