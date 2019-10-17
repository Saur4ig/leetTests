package _4_Longest_Common_Prefix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "Example 1",
			input: []string{"flower", "flow", "flight"},
			want:  "fl",
		},
		{
			name:  "Example 2",
			input: []string{"dog", "racecar", "car"},
			want:  "",
		},
		{
			name:  "Example 3",
			input: []string{"", "b"},
			want:  "",
		},
		{
			name:  "Example 4",
			input: []string{"a", "b", ""},
			want:  "",
		},
		{
			name:  "Example 5",
			input: []string{"a", "ac"},
			want:  "a",
		},
		{
			name:  "Example 6",
			input: []string{},
			want:  "",
		},
		{
			name:  "Example 7",
			input: []string{"a"},
			want:  "a",
		},
		{
			name:  "Example 8",
			input: []string{"aa", "a"},
			want:  "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
