package _25_Valid_Palindrome

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Example 1",
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
		{
			name:  "Example 2",
			input: "race a car",
			want:  false,
		},
		{
			name:  "Example 3",
			input: "0P",
			want:  false,
		},
		{
			name:  "Example 4",
			input: "a",
			want:  true,
		},
		{
			name:  "Example 5",
			input: "",
			want:  true,
		},
		{
			name:  "Example 6",
			input: " ",
			want:  true,
		},
		{
			name:  "Example 6",
			input: "a.",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
