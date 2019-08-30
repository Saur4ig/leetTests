package numbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "Example 1",
			input: 123,
			want:  321,
		},
		{
			name:  "Example 2",
			input: -123,
			want:  -321,
		},
		{
			name:  "Example 3",
			input: 120,
			want:  21,
		},
		{
			name:  "Example 4",
			input: 1,
			want:  1,
		},
		{
			name:  "Example 5",
			input: 0,
			want:  0,
		},
		{
			name:  "Example 6",
			input: 100,
			want:  1,
		},
		{
			name:  "Example 7",
			input: 1534236469,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
