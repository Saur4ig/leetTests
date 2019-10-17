package _8_strStr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		needle string
		want   int
	}{
		{
			name:   "Example 1",
			input:  "hello",
			needle: "ll",
			want:   2,
		},
		{
			name:   "Example 2",
			input:  "aaaaa",
			needle: "bba",
			want:   -1,
		},
		{
			name:   "Example 3",
			input:  "aaaaa",
			needle: "aa",
			want:   0,
		},
		{
			name:   "Example 4",
			input:  "aabaab",
			needle: "aab",
			want:   0,
		},
		{
			name:   "Example 5",
			input:  "",
			needle: "",
			want:   0,
		},
		{
			name:   "Example 6",
			input:  "",
			needle: "aa",
			want:   -1,
		},
		{
			name:   "Example 7",
			input:  "a",
			needle: "",
			want:   0,
		},
		{
			name:   "Example 8",
			input:  "mississippi",
			needle: "issip",
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strStr(tt.input, tt.needle)
			require.Equal(t, tt.want, got)
		})
	}
}
