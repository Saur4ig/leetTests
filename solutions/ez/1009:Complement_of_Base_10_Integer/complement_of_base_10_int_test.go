package _009_Complement_of_Base_10_Integer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBitwiseComplement(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "Example 1",
			input: 5,
			want:  2,
		},
		{
			name:  "Example 2",
			input: 7,
			want:  0,
		},
		{
			name:  "Example 3",
			input: 10,
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bitwiseComplement(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
