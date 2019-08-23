package numbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example 1",
			input: "III",
			want:  3,
		},
		{
			name:  "Example 2",
			input: "IV",
			want:  4,
		},
		{
			name:  "Example 3",
			input: "IX",
			want:  9,
		},
		{
			name:  "Example 4",
			input: "LVIII",
			want:  58,
		},
		{
			name:  "Example 5",
			input: "MCMXCIV",
			want:  1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
