package stringsE

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Example 1",
			input: "()",
			want:  true,
		},
		{
			name:  "Example 2",
			input: "[]",
			want:  true,
		},
		{
			name:  "Example 3",
			input: "{}",
			want:  true,
		},
		{
			name:  "Example 4",
			input: "{(})",
			want:  false,
		},
		{
			name:  "Example 5",
			input: "((]]",
			want:  false,
		},
		{
			name:  "Example 6",
			input: "((]]",
			want:  false,
		},
		{
			name:  "Example 7",
			input: ")(",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
