package stringsE

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetectCapitalUse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Example 1",
			input: "USA",
			want:  true,
		},
		{
			name:  "Example 2",
			input: "leetcode",
			want:  true,
		},
		{
			name:  "Example 3",
			input: "Google",
			want:  true,
		},
		{
			name:  "Example 4",
			input: "FlaG",
			want:  false,
		},
		{
			name:  "Example 5",
			input: "tEstGs",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectCapitalUse(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
