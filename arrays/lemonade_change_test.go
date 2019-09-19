package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLemonadeChange(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "Example 1",
			input: []int{5, 5, 5, 10, 20},
			want:  true,
		},
		{
			name:  "Example 2",
			input: []int{5, 5, 10},
			want:  true,
		},
		{
			name:  "Example 3",
			input: []int{10, 10},
			want:  false,
		},
		{
			name:  "Example 4",
			input: []int{5, 5, 10, 10, 20},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lemonadeChange(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
